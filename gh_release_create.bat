gl2gh.exe
setlocal
set M="initial release"
call gi %M%
gh release create v1.0.0 --target main --notes %M%
