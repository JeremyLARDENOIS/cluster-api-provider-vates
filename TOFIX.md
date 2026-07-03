# Code Review — À corriger

## Critique

### vatescluster_controller.go

- **Ready = true avant découverte endpoint** (l.110-112) — CAPI pense que l'infra est prête alors que le control plane endpoint est encore vide.
- **`r.Update()` appelé à chaque réco** pour le finalizer (l.65-69), sans vérifier s'il a vraiment changé.
- **Pas de re-fetch avant `Status().Update()`** (l.114-117) — risque de conflit `object has been modified`.
- **`getOwnerCluster()` liste tous les Clusters du namespace** (l.150-165) — utiliser `OwnerReferences` à la place.
- **Ne watch pas les Clusters/Machines** (l.210-215) — changements sur le Cluster parent ne déclenchent pas de réco.

### vm_ops.go

- **ProviderID dans Spec + Status sans re-fetch entre les deux** (l.137-148) — conflit potentiel.
- **Erreurs de résolution réseau silencieuses** (l.59-62) — `continue` sans erreur, VM créée avec moins de VIFs que demandé.
- **Orphelin XO si `c.Update()` ou `Status().Update()` échoue** après `CreateVM` (l.137-148) — VM existe dans XO mais pas dans K8s.

### cloud_config.go

- **`GetOwnerMachine()` liste TOUTES les Machines du namespace** (l.183-196) — O(n) à chaque réco, utiliser `OwnerReferences`.
- **`GetVatesCluster` pas de nil-check sur `InfrastructureRef`** (l.161-166).

### resolve.go

- **`ResolveTemplateID` et `ResolvePoolID` pas de nil-check sur `V1Client()`** (l.72, 100) — panic si nil.

### vatesmachine_reconcile.go

- **`Start() + CleanShutdown()` en séquence sans attendre** (l.247-248) — race condition dans la suppression de VM.

### Tests

- **`vatesmachine_controller_test.go` / `vatescluster_controller_test.go`** — tests scaffold vides avec `TODO(user)`, sans assertions utiles.
- **Pas de tests unitaires pour `CreateVM`, `StartVM`, `WaitForVMReady`** — fonctions critiques non couvertes.
- **Chemins d'erreur de `reconcileDelete` non testés** — échecs Start/CleanShutdown/Delete.

---

## Moyen

| Fichier | Problème |
|---------|----------|
| `api/v1beta2/vatescluster_types.go:10` | `KubeVIPSpec.Enabled` sans `omitempty`, toujours requis |
| `api/v1beta2/vatesmachine_types.go:83-87` | `Network.Name` casse le pattern `NetworkID` / `NetworkName` |
| `api/v1beta2/vatesmachine_types.go:142` | Print column `Ready` en `type=string` au lieu de `type=boolean` |
| `vatescluster_controller.go:225` / `status.go:29` | `LastTransitionTime` mis à jour à chaque condition write (devrait être seulement quand le status change) |
| `status.go:33-38` | `ObservedGeneration` jamais défini dans les conditions |
| `vatescluster_controller.go:139` | Log "Cleaning up" mais rien n'est nettoyé |
| `vatescluster_controller.go:187-193` | Erreur de `Get` silencieuse dans `discoverControlPlaneEndpoint` |
| `client.go:17-18` | `GetOrCreateXOClient` retourne nil, nil sans condition → machine "pending" sans explication |
| `vm_ops.go:85-105` | `DiskSize` invalide silencieusement ignoré |
| `resolve.go:119` | `ResolveNetworkID` n'a pas de paramètre `ctx` (incohérent avec les autres) |

---

## Léger / style

- `vm_ops.go:199,289` — variable `vm` shadowée deux fois
- `cloud_config.go:66-81` — `ssh_authorized_keys` lu deux fois de la map
- `cloud_config.go:88-90` — nil-map guard après les lectures (marche par accident)
- `vm_ops.go:253` — adresse de variable de boucle passée à `ConnectVIF`
- `resolve.go:63,91` — paramètre `ctx` inutilisé dans `ResolveTemplateID`/`ResolvePoolID`
- Messages de log au présent au lieu du passé (vm_ops.go:40)
- `PROJECT` — manque l'entrée `VatesMachineTemplate`
