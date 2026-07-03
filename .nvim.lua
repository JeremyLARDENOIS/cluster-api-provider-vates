local root = vim.fn.getcwd()
local custom_lint = root .. "/bin/golangci-lint"

if vim.fn.filereadable(custom_lint) == 1 then
	local ok, lint = pcall(require, "lint")
	if ok and lint.linters.golangcilint then
		lint.linters.golangcilint.cmd = custom_lint
	end

	vim.env.PATH = root .. "/bin:" .. vim.env.PATH
end
