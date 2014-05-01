## Go by Example 中文翻译

本翻译除了为使用 github page 来架修改了部分外，完全遵从原作的模式。现已经完成了所有文件的初步翻译，比较粗糙，其中纰漏也很多，欢迎 fork 并提交 pull request。

还等什么，[马上开始学习吧](http://everyx.github.io/gobyexample/)

### 为了 gh-pages 而做的一些修改

```diff
diff --git a/tools/generate.go b/tools/generate.go
index 30d7291..e1caa03 100644
--- a/tools/generate.go
+++ b/tools/generate.go
@@ -15,7 +15,7 @@ import (
 )

 var cacheDir = "/tmp/gobyexample-cache"
-var siteDir = "./public"
+var siteDir = "./"
 var pygmentizeBin = "./vendor/pygments/pygmentize"

 func check(err error) {
@@ -221,7 +221,17 @@ func parseExamples() []*Example {
     examples := make([]*Example, 0)
     for _, exampleName := range exampleNames {
         if (exampleName != "") && !strings.HasPrefix(exampleName, "#") {
-            example := Example{Name: exampleName}
+            exampleNameZh := exampleName
+            if strings.Index(exampleName, "->") != -1 {
+                names := strings.Split(exampleName, "->")
+                exampleName = names[0]
+                if strings.Trim(names[1], " ") != "" {
+                    exampleNameZh = names[1]
+                } else {
+                    exampleNameZh = names[0]
+                }
+            }
+            example := Example{Name: exampleNameZh}
             exampleId := strings.ToLower(exampleName)
             exampleId = strings.Replace(exampleId, " ", "-", -1)
             exampleId = strings.Replace(exampleId, "/", "-", -1)
@@ -270,13 +280,14 @@ func renderExamples(examples []*Example) {
     _, err := exampleTmpl.Parse(mustReadFile("templates/example.tmpl"))
     check(err)
     for _, example := range examples {
-        exampleF, err := os.Create(siteDir + "/" + example.Id)
+        exampleF, err := os.Create(siteDir + "/" + example.Id + ".html")
         check(err)
         exampleTmpl.Execute(exampleF, example)
     }
 }

 func main() {
+    ensureDir(siteDir)
     copyFile("templates/site.css", siteDir+"/site.css")
     copyFile("templates/favicon.ico", siteDir+"/favicon.ico")
     copyFile("templates/404.html", siteDir+"/404.html")
```

### License

This work is copyright Mark McGranaghan and licensed under a
[Creative Commons Attribution 3.0 Unported License](http://creativecommons.org/licenses/by/3.0/).

The Go Gopher is copyright the Go Authors and licensed under a
[Creative Commons Attribution 3.0 Unported License](http://creativecommons.org/licenses/by/3.0/).

### 原版 ReadMe

[ReadMe] (https://github.com/mmcgrana/gobyexample#go-by-example) 
