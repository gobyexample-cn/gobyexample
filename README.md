## Go by Example 中文翻译

[原版 ReadMe](https://github.com/mmcgrana/gobyexample#go-by-example)

[马上开始学习吧](http://everyx.github.io/gobyexample/)

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

### 翻译进度

|内容|完成|
|:-----------------------------|:--|
|Hello World |✔|
|Values |✔| 
|Variables |✔|
|Constants |✔|
|For |✔|
|If/Else |✔|
|Switch |✔|
|Arrays |✔|
|Slices |✔|
|Maps |✔|
|Range |✔|
|Functions |✔|
|Multiple Return Values|✔|
|Multiple Return Values|✔|
|Variadic Functions|✔|
|Closures|✔|
|Recursion|✔|
|Pointers|✔|
|Structs|✔|
|Methods|✔|
|Interfaces|✔|
|Errors|✔|
|Goroutines|✔|
|Channels|✔|
|Channel Buffering||
|Channel Synchronization||
|Channel Directions||
|Select||
|Timeouts||
|Non-Blocking Channel Operations||
|Closing Channels||
|Range over Channels||
|Timers||
|Tickers||
|Worker Pools||
|Rate Limiting||
|Atomic Counters||
|Mutexes||
|Stateful Goroutines||
|Sorting||
|Sorting by Functions||
|Panic||
|Defer||
|Collection Functions||
|String Functions||
|String Formatting||
|Regular Expressions||
|JSON||
|Time||
|Epoch||
|Time Formatting / Parsing||
|Random Numbers||
|Number Parsing||
|URL Parsing||
|SHA1 Hashes||
|Base64 Encoding||
|Reading Files||
|Writing Files||
|Line Filters||
|Command-Line Arguments||
|Command-Line Flags||
|Environment Variables||
|Spawning Processes||
|Exec'ing Processes||
|Signals||
|Exit||

### License

This work is copyright Mark McGranaghan and licensed under a
[Creative Commons Attribution 3.0 Unported License](http://creativecommons.org/licenses/by/3.0/).

The Go Gopher is copyright the Go Authors and licensed under a
[Creative Commons Attribution 3.0 Unported License](http://creativecommons.org/licenses/by/3.0/).
