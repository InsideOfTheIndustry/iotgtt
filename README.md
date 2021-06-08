# Some test and benchmark 

Benchmark for understanding golang.
## notice 
If you use VSCODE for coding, you need to config your gopls in vscode's settings.json
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