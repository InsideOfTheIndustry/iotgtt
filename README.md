# Some test and benchmark 

Benchmark for understanding golang.
## notice 
If you use `VSCODE` for coding, you need to config your gopls in vscode's settings.json for support multi module file.
Add code like below:

```json
  "gopls": {
  "build.experimentalWorkspaceModule": true
  }
```
## append & copy

Copy is more effective than append.

## calcu md5 of a file 
slice the file is more effective than read the whole file to memeroy.

## rand seed 


## string to []byte 
Use pointer is much more faster. 
使用反射和指针性能提升十分明显

# Memory escape
用使用指针一定比内存快？