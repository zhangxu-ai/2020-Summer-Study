# HTML笔记
## HTML简介
超文本标记语言（英语：HyperText Markup Language，简称：HTML）是一种用于创建网页的标准标记语言。可以使用 HTML 来建立自己的 WEB 站点，HTML 运行在浏览器上，由浏览器来解析。
html的后缀名：
* .html
* .htm
以上两种后缀名等效，但推荐使用第一种。

示例代码如下：
```html
<!DOCTYPE html>
<html>
<head>
<meta charset="utf-8">
<title>AHUTSJ_HMTL</title>
</head>
<body>
    <h1>我的第一个标题</h1>
    <p>我的第一个段落。</p>
</body>
</html>
```
### 实例解析：
* DOCTYPE html 声明为HTML5文档
* html 元素是HTML文档的根元素
* head 元素包含了文档的meta数据，如网页编码方式等
* title 元素定义了文档的标题
* body 元素包含了文档的可见页面内容
* h1 元素定义了一个大标题
* p 元素定义了一个段落

### HTML标签
HTML 标记标签通常被称为 HTML 标签 (HTML tag)。
* HTML标签是由尖括号包含的关键词
* HTML标签通常成对出现，分别表示开始和结束

## HTML基础——4个实例
### HTML标题
HTML 标题（Heading）是通过 h1-h6 标签来定义的，类似于MD的六级标题。
```html
<!DOCTYPE html>
<html>
<head>
<meta charset="utf-8">
<title>AHUTSJ_HTML</title>
</head>
<body>

<h1>这是标题 1</h1>
<h2>这是标题 2</h2>
<h3>这是标题 3</h3>
<h4>这是标题 4</h4>
<h5>这是标题 5</h5>
<h6>这是标题 6</h6>

</body>
</html>
```
### HTML段落
HTML 段落是通过标签 p 来定义的。
```html
<!DOCTYPE html>
<html>
<head>
<meta charset="utf-8">
<title>菜鸟教程(runoob.com)</title>
</head>
<body>

<p>这是一个段落。</p>
<p>这是一个段落。</p>
<p>这是一个段落。</p>

</body>
</html>
```
### HTML链接
HTML 链接是通过标签 a 来定义的。
```html
<!DOCTYPE html>
<html>
<head>
<meta charset="utf-8">
<title>AHUTSJ_HTML</title>
</head>
<body>

<a href="https://www.baidu.com">这是一个链接使用了 href 属性</a>

</body>
</html>
```
**提示**：在 href 属性中指定链接的地址。
### HTML图像
HTML 图像是通过标签 img 来定义的。
```html
<!DOCTYPE html>
<html>
<head>
<meta charset="utf-8">
<title>AHUTSJ_HTML</title>
</head>
<body>

<img src="/images/logo.png" width="258" height="39" />

</body>
</html>
```
**注意**：图像的名称和尺寸是以属性的形式提供的。
## HTML元素
### HTML元素语法
* HTML元素以**开始标签**起始
* HTML元素以**结束标签**终止
* **元素的内容**是开始标签和结束标签之间的内容
* 某些HTML元素具有空内容
* 空元素**在开始标签中进行关闭**（以开始标签的结束而结束）
* 大多数HTML元素可拥有**属性**
### 嵌套的HTML元素
大多数HTML元素可以嵌套（HTML元素可以包含其他HTML元素）。HMTL文档由相互嵌套的HTML元素构成。
HMLT文档实例：
```html
<!DOCTYPE html>
<html>

<body>
<p>这是第一个段落。</p>
</body>

</html>
```
以上实例包含了三个HTML元素。
**注意**：不要忘记结束标签！
### HTML空元素
没有内容的 HTML 元素被称为空元素。空元素是在开始标签中关闭的。在开始标签中添加斜杠，是关闭空元素的正确方法。
### HTML提示：使用小写标签
HTML标签对大小写不敏感，但是推荐使用小写标签。
## HTML属性
属性是HTML元素提供的附加信息，具有以下特点：
* HTML元素可以设置属性
* 属性可以再元素中添加附加信息
* 属性一般描述于**开始标签**
* 属性总是以名称/值对的形式出现，**比如：name="value"**
### 属性实例
HTML链接由 a 标签定义，链接的地址在**herf属性**中指定：
```html
<!DOCTYPE html>
<html>
<head>
<meta charset="utf-8">
<title>AHUTSJ_HTML</title>
</head>
<body>

<a href="https://www.baidu.com">这是一个链接使用了 href 属性</a>

</body>
</html>
```
### HTML属性常用引用属性值
属性值应该始终被包含在引号内，双引号是最常用的，但是使用单引号也没问题。
**提示**：在某些个别情况下，比如属性值本身就有双引号，那必须使用单引号。
## HTML水平线
hr 标签在HTML页面中创建水平线。
hr元素可用于分割内容。
```html
<!DOCTYPE html>
<html>
<head>
<meta charset="utf-8">
<title>AHUTSJ_HTML</title>
</head>
<body>
	<p>hr 标签定义水平线：</p>
	<hr />
	<p>这是段落。</p>
	<hr />
	<p>这是段落。</p>
	<hr />
	<p>这是段落。</p>
</body>
</html>
```
## HTML注释
可以将注释插入 HTML 代码中，这样可以提高其可读性，使代码更易被人理解。浏览器会忽略注释，也不会显示它们。注释写法如下：
```html
<!DOCTYPE html>
<html>
<head>
<meta charset="utf-8">
<title>AHUTSJ_HTML</title>
</head>
<body>

<!--这是一个注释，注释在浏览器中不会显示-->

<p>这是一个段落</p>

</body>
</html>
```
**注释**：开始括号之后（左边的括号）需要紧跟一个叹号，结束括号之前（右边的括号）不需要，合理地使用注释可以对未来的代码编辑工作产生帮助。
## HTML拆行
如果希望在不产生新段落的情况下进行换行，使用 br 标签：
```html
<!DOCTYPE html>
<html>
<head>
<meta charset="utf-8">
<title>AHUTSJ_HTML</title>
</head>
<body>

<p>这个<br>段落<br>演示了分行的效果</p>

</body>
</html>
```
**提示**：br 元素是一个空的HTML元素，由于关闭标签没有任何意义，因此没有结束标签。
## HTML文本格式化
```html
<!DOCTYPE html>
<html>
<head> 
<meta charset="utf-8"> 
<title>AHUTSJ_HTML</title> 
</head> 
<body>

<b>加粗文本</b><br><br>
<i>斜体文本</i><br><br>
<code>电脑自动输出</code><br><br>
这是 <sub> 下标</sub> 和 <sup> 上标</sup>

</body>
</html>
```
### HTML格式化标签
HTML使用标签 b 和 i 对输出的文本进行格式化，这些HTML标签被称为格式化标签。
**提示**：通常标签 strong 替换标签 b ，em 替换 i。
格式化实例：
1. 文本格式化：
```html
<!DOCTYPE html>
<html>
<head>
<meta charset="utf-8">
<title>AHUTSJ_HTML</title>
</head>

<body>

<b>这个文本是加粗的</b>
<br />

<strong>这个文本是加粗的</strong>
<br />

<big>这个文本字体放大</big>
<br />

<em>这个文本是斜体的</em>
<br />

<i>这个文本是斜体的</i>
<br />

<small>这个文本是缩小的</small>
<br />

这个文本包含
<sub>下标</sub>
<br />

这个文本包含
<sup>上标</sup>

</body>
</html>
```
2. 预格式文本：
```html
<!DOCTYPE html>
<html>
<head> 
<meta charset="utf-8"> 
<title>AHUTSJ_HTML</title> 
</head>
<body>

<pre>
此例演示如何使用 pre 标签
对空行和    空格
进行控制
</pre>

</body>
</html>
```
3. "计算机输出"标签：
```html
<!DOCTYPE html> 
<html>
<head> 
<meta charset="utf-8"> 
<title>AHUTSJ_HTML</title> 
</head>
<body>

<code>计算机输出</code>
<br />
<kbd>键盘输入</kbd>
<br />
<tt>打字机文本</tt>
<br />
<samp>计算机代码样本</samp>
<br />
<var>计算机变量</var>
<br />

<p>
<b>注释：</b>这些标签常用于显示计算机/编程代码。
</p>

</body>
</html>
```
4. 地址:
```html
<!DOCTYPE html>
<html>
<head> 
<meta charset="utf-8"> 
<title>菜鸟教程(runoob.com)</title> 
</head>
<body>

<address>
Written by <a href="mailto:webmaster@example.com">Jon Doe</a>.<br> 
Visit us at:<br>
Example.com<br>
Box 564, Disneyland<br>
USA
</address>

</body>
</html>
```
5. 缩进和首字母缩写
此例演示如何实现缩写或首字母缩写。
```html
<!DOCTYPE html> 
<html>
<head> 
<meta charset="utf-8"> 
<title>AHUTSJ_HTML</title> 
</head>
<body>

<abbr title="etcetera">etc.</abbr>
<br />
<acronym title="World Wide Web">WWW</acronym>

<p>在某些浏览器中，当您把鼠标移至缩略词语上时，title 可用于展示表达的完整版本。</p>

<p>仅对于 IE 5 中的 acronym 元素有效。</p>

<p>对于 Netscape 6.2 中的 abbr 和 acronym 元素都有效。</p>

</body>
</html>
```
6. 文字方向
此例演示文字的显示方向。
```html
<!DOCTYPE html>
<html>
<head> 
<meta charset="utf-8"> 
<title>AHUTSJ_HTML</title> 
</head> 
<body>

<p>该段落文字从左到右显示。</p>  
<p><bdo dir="rtl">该段落文字从右到左显示。</bdo></p>  

</body>
</html>
```
7. 块引用
此例演示如何实现长短不一的引用语
```html
<!DOCTYPE html>
<html>
<head> 
<meta charset="utf-8"> 
<title>AHUTSJ_HTML</title> 
</head>
<body>

<p>WWF's goal is to: 
<q>Build a future where people live in harmony with nature.</q>
We hope they succeed.</p>

</body>
</html>
```
8. 删除字效果和插入字效果
此例演示如何标记删除文本和插入文本
```html
<!DOCTYPE html>
<html>
<head> 
<meta charset="utf-8"> 
<title>AHUTSJ_HTML</title> 
</head>
<body>

<p>My favorite color is <del>blue</del> <ins>red</ins>!</p>

</body>
</html>
```
## HTML链接
HTML 使用超级链接与网络上的另一个文档相连。几乎可以在所有的网页中找到链接。点击链接可以从一张页面跳转到另一张页面。
```html
<!DOCTYPE html>
<html>
<head> 
<meta charset="utf-8"> 
<title>AHUTSJ_HTML</title> 
</head>
<body>

<p><a href="//www.microsoft.com/">本文本</a> 是一个指向万维网上的页面的链接。</p>

</body>
</html>
```
HTML使用标签 a 来设置超文本链接。超链接可以是一个字，一个词，或者一组词，也可以是一幅图像，您可以点击这些内容来跳转到新的文档或者当前文档中的某个部分。当您把鼠标指针移动到网页中的某个链接上时，箭头会变为一只小手。在标签 a 中使用了 href 属性来描述链接的地址。
默认情况下，链接将以以下形式出现在浏览器中：
* 一个未访问过的链接显示为蓝色字体并带有下划线
* 访问过的链接显示为紫色并带有下划线
* 点击链接时，链接显示为红色并带有下划线
**注意**：如果为这些链接设置了CSS样式，展示样式会根据CSS的设定来显示。
### HTML链接语法
链接的HTML代码很简单：
`<a href="URL">链接文本</a>`
其中 href 属性描述了链接的目标。
**提示**：“链接文本”不一定是文本，图片或者其他HTML元素都可以成为链接。
### HTML链接——target属性
使用 target 属性，你可以定义被链接的文档在何处显示。
```html
<!DOCTYPE html>
<html>
<head> 
<meta charset="utf-8"> 
<title>AHUTSJ_HTML</title> 
</head>
<body>

<a href="https://www.baidu.com/" target="_blank">访问菜鸟教程!</a>

<p>如果你将 target 属性设置为 &quot;_blank&quot;, 链接将在新窗口打开。</p>

</body>
</html>
```
### HTML链接——id属性
id属性可用于创建一个在HTML文档书签标记。
**提示**：书签是不以任何特殊的方式显示，在HTML文档中是不显示的，所以对于读者来说是隐藏的。
#### id属性实例
在HTML文档中插入id：
`<a id="tips">小提示</a>`
在HTML文档中创建一个链接到“小提示”（id=tips）:
`<a href="#tips">访问小提示</a>`
或者，从另一个页面创建一个链接到“小提示”（id=tips）:
`<a href="https://www.baidu.com">访问小提示</a>`
#### 从当前界面跳转到另一个界面
```html
<!DOCTYPE html>
<html>
<head>
<meta charset="utf-8">
<title>AHUTSJ_HTML</title>
</head>
<body>

<p>
<a href="#C4">查看章节 4</a>
</p>

<h2>章节 1</h2>
<p>这边显示该章节的内容……</p>

<h2>章节 2</h2>
<p>这边显示该章节的内容……</p>

<h2>章节 3</h2>
<p>这边显示该章节的内容……</p>

<h2><a id="C4">章节 4</a></h2>
<p>这边显示该章节的内容……</p>

<h2>章节 5</h2>
<p>这边显示该章节的内容……</p>

</body>
</html>
```
## HTML头部
1. title 定义了HTML文档的标题：
```html
<!DOCTYPE html>
<html>
<head>
<meta charset="utf-8"> 
<title>我的 HTML 的第一页</title>
</head>

<body>
<p>浏览器中包含body元素的内容。</p>
<p>浏览器的标题包含title元素的内容</p>
</body>

</html>
```
2. base 定义了所有链接的URL，使用 base 定义了页面内所以链接默认的链接目标地址。
3. meta 提供了HTML文档的meta标记
使用 meta 元素描述HTML文档的描述，关键词、作者、字符集等。
```html
<!DOCTYPE html>
<html>
<head>
<meta charset="utf-8"> 
<title>AHUTSJ_HTML</title> 
<meta name="description" content="HTML学习">
<meta name="keywords" content="HTML">
<meta name="AHUTSJ" content="HTML_LEARNING">
</head>
<body>

<p>所有 meta 标签显示在 head 部分</p>

</body>
</html>
```
### HTML head 元素
head 元素包含了所有的头部标签元素，在 head 元素中可以插入脚本（scripts）, 样式文件（CSS），及各种meta信息。
可以添加在头部区域的元素标签为:title、style、meta、link、script、noscript和base。
## HTML样式-CSS
CSS用于渲染HTML元素标签的样式。
```html
<!DOCTYPE html>
<html>
<head> 
<meta charset="utf-8"> 
<title>AHUTSJ_HTML</title> 
</head>
<body>

<div style="opacity:0.5;position:absolute;left:50px;width:300px;height:150px;background-color:#40B3DF"></div>

<div style="font-family:verdana;padding:20px;border-radius:10px;border:10px solid #EE872A;">

<div style="opacity:0.3;position:absolute;left:120px;width:100px;height:200px;background-color:#8AC007"></div>

<h3>Look! Styles and colors</h3>

<div style="letter-spacing:12px;">Manipulate Text</div>

<div style="color:#40B3DF;">Colors
<span style="background-color:#B4009E;color:#ffffff;">Boxes</span>
</div>

<div style="color:#000000;">and more...</div>

</div>

</body>
</html>
```
#### 添加 head 部分的信息对HTML格式化
```html
<!DOCTYPE html>
<html>
<head>
<meta charset="utf-8"> 
<title>AHUTSJ_HTML</title>
<style type="text/css">
h1 {color:red;}
p {color:blue;}
</style>
</head>

<body>
<h1>这是一个标题</h1>
<p>这是一个段落。</p>
</body>

</html>
``` 
#### 没有下划线的链接
```html
<!DOCTYPE html>
<html>
<head> 
<meta charset="utf-8"> 
<title>AHUTSJ_HTML</title> 
</head>
<body>

<a href="//www.baidu.com/" style="text-decoration:none;">访问 baidu.com!</a>
<!--style="text-decoration:none;"表示去掉链接下默认的下划线-->
</body>
</html>
```
### 如何使用CSS
CSS是在HTML4开始使用的，是为了更好的渲染HTML元素引用的。CSS可以通过以下方式添加到HTML中：
* 内部样式——在HTML元素使用style属性
* 内部样式表——在HTML文档头部 head 区域使用style元素来包含CSS
* 外部引用——使用外部CSS文件
#### 内联样式
当特殊的样式需要应用到个别元素时，就可以使用内联样式。 使用内联样式的方法是在相关的标签中使用样式属性。样式属性可以包含任何 CSS 属性。以下实例显示出如何改变段落的颜色和左外边距。
`<p style="color:blue;margin-left:20px;">这是一个段落。</p>`
#### HTML样式实例——背景颜色
背景色属性（background-color）定义一个元素的背景颜色：
```html
<!DOCTYPE html>
<html>
<head> 
<meta charset="utf-8"> 
<title>AHUTSJ_HTML</title> 
</head>
<body style="background-color:yellow;">
<h2 style="background-color:red;">这是一个标题</h2>
<p style="background-color:green;">这是一个段落。</p>
</body>
</html>
```
#### HTML 样式实例 - 字体, 字体颜色 ，字体大小
我们可以使用font-family（字体），color（颜色），和font-size（字体大小）属性来定义字体的样式:
```html
<!DOCTYPE html>
<html>
<head> 
<meta charset="utf-8"> 
<title>AHUTSJ_HTML</title> 
</head>
<body>
<h1 style="font-family:verdana;">一个标题</h1>
<p style="font-family:arial;color:red;font-size:20px;">一个段落。</p>
</body>
</html>
```
#### 内部样式表
当单个文件需要特别样式时，就可以使用内部样式表。你可以在 head  部分通过 style 标签定义内部样式表:
```html
<head>
<style type="text/css">
body {background-color:yellow;}
p {color:blue;}
</style>
</head>
```
#### 外部样式表
当样式需要被应用到很多页面的时候，外部样式表将是理想的选择。使用外部样式表，你就可以通过更改一个文件来改变整个站点的外观。
```html
<head>
<link rel="stylesheet" type="text/css" href="mystyle.css">
</head>
```
### 已弃用的标签和属性
在HTML 4, 原来支持定义HTML元素样式的标签和属性已被弃用。这些标签将不支持新版本的HTML标签。
不建议使用的标签有: font , center , strike
不建议使用的属性: color 和 bgcolor.
## HTML图像
```html
<!DOCTYPE html>
<html>
<head> 
<meta charset="utf-8"> 
<title>AHUTSJ_HTML</title> 
</head>
<body>

<p>
一个图像:
<img src="smiley.gif" alt="Smiley face" width="32" height="32"></p>

<p>
一个动图:
<img src="hackanm.gif" alt="Computer man" width="48" height="48"></p>

<p>
注意插入动图的语法和静态图的语法是一样的。
</p>

</body>
</html>
```
### HTML图像标签（img）和源属性（src）
在HTML中，图像由 img 标签定义。 img 是空标签，只包含属性，并没有闭合标签。
要在页面上显示图像，需要使用源属性（src），源属性的值是图像的URL地址。
定义图像的语法是：
`<img src="URL" alt="some_text">`
URL是存储图像的位置。浏览器将图像显示在文档中图像标签出现的地方。如果你将图像标签置于两个段落之间，那么浏览器会首先显示第一个段落，然后显示图片，最后显示第二段。
### HTML图像——Alt属性
alt 属性用来为图像定义一串预备的可替换的文本。替换的文本是用户自定义的。
`<img src="boat.gif" alt="Big Boat">`
在浏览器无法载入图像时，替换文本属性告诉读者她们失去的信息。此时，浏览器将显示这个替代性的文本而不是图像。为页面上的图像都加上替换文本属性是个好习惯，这样有助于更好的显示信息，并且对于那些使用纯文本浏览器的人来说是非常有用的。
### HTML图像——设置图像的高度和宽度
height（高度） 与 width（宽度）属性用于设置图像的高度与宽度。属性的默认值是像素：
`<img src="pulpit.jpg" alt="Pulpit rock" width="304" height="228">`
**提示**：指定图像的高度和宽度是一个很好的习惯。如果图像指定了高度宽度，页面加载时就会保留指定的尺寸。如果没有指定图片的大小，加载页面时有可能会破坏HTML页面的整体布局。
### 实例——创建图片映射

```html
<!DOCTYPE html>
<html>
<head>
<meta charset="utf-8">
<title>AHUTSJ_HTML</title>
</head>
<body>
​
<p>点击太阳或其他行星，注意变化：</p>
​
<img src="planets.gif" width="145" height="126" alt="Planets" usemap="#planetmap">
​
<map name="planetmap">
  <area shape="rect" coords="0,0,82,126" alt="Sun" href="sun.htm">
  <area shape="circle" coords="90,58,3" alt="Mercury" href="mercur.htm">
  <area shape="circle" coords="124,58,8" alt="Venus" href="venus.htm">
</map>
​
</body>
</html>
```
## HTML表格
HTML表格由 table 来定义，每个表格均有若干行（由 tr 标签来定义），每行被分割为若干单元格（由 td 标签来定义）。
```html
<!DOCTYPE html>
<html>
<head> 
<meta charset="utf-8"> 
<title>AHUTSJ_HTML</title> 
</head>
<body>

<p>
每个表格从一个 table 标签开始。 
每个表格行从 tr 标签开始。
每个表格的数据从 td 标签开始。
</p>

<h4>一列:</h4>
<table border="1">
<tr>
  <td>100</td>
</tr>
</table>

<h4>一行三列:</h4>
<table border="1">
<tr>
  <td>100</td>
  <td>200</td>
  <td>300</td>
</tr>
</table>

<h4>两行三列:</h4>
<table border="1">
<tr>
  <td>100</td>
  <td>200</td>
  <td>300</td>
</tr>
<tr>
  <td>400</td>
  <td>500</td>
  <td>600</td>
</tr>
</table>

</body>
</html>
```
### HTML表格表头
表格的表头由 th 标签进行定义。大多数浏览器会把表头显示为粗体居中的文本：
```html
<table border="1">
    <tr>
        <th>Header 1</th>
        <th>Header 2</th>
    </tr>
    <tr>
        <td>row 1, cell 1</td>
        <td>row 1, cell 2</td>
    </tr>
    <tr>
        <td>row 2, cell 1</td>
        <td>row 2, cell 2</td>
    </tr>
</table>
```
### 没有边框的表格
```html
<table><!--border的默认值是0，表示无边框-->
<tr>
  <td>100</td>
  <td>200</td>
  <td>300</td>
</tr>
<tr>
  <td>400</td>
  <td>500</td>
  <td>600</td>
</tr>
</table>

<h4>这个表格没有边框:</h4>
<table border="0">
<tr>
  <td>100</td>
  <td>200</td>
  <td>300</td>
</tr>
<tr>
  <td>400</td>
  <td>500</td>
  <td>600</td>
</tr>
</table>

</body>
</html>
```
### 表格中的表头
```html
<!DOCTYPE html>
<html>
<head> 
<meta charset="utf-8"> 
<title>AHUTSJ_HTML</title> 
</head>
<body>

<h4>水平标题:</h4>
<table border="1">
<tr>
  <th>Name</th>
  <th>Telephone</th>
  <th>Telephone</th>
</tr>
<tr>
  <td>Bill Gates</td>
  <td>555 77 854</td>
  <td>555 77 855</td>
</tr>
</table>

<h4>垂直标题:</h4>
<table border="1">
<tr>
  <th>First Name:</th>
  <td>Bill Gates</td>
</tr>
<tr>
  <th>Telephone:</th>
  <td>555 77 854</td>
</tr>
<tr>
  <th>Telephone:</th>
  <td>555 77 855</td>
</tr>
</table>

</body>
</html>
```
### 带有标题的表格
```html
<!DOCTYPE html>
<html>
<head> 
<meta charset="utf-8"> 
<title>AHUTSJ_HTML</title> 
</head>
<body>

<table border="1">
  <caption>Monthly savings</caption><!-- caption 标签用来表示一个表格的标题-->
  <tr>
    <th>Month</th>
    <th>Savings</th>
  </tr>
  <tr>
    <td>January</td>
    <td>$100</td>
  </tr>
  <tr>
    <td>February</td>
    <td>$50</td>
  </tr>
</table>

</body>
</html>
```
### 跨行跨列的表格
```html
<!DOCTYPE html>
<html>
<head> 
<meta charset="utf-8"> 
<title>AHUTSJ_HTML</title> 
</head>
<body>

<h4>单元格跨两列:</h4>
<table border="1">
<tr>
  <th>Name</th>
  <th colspan="2">Telephone</th><!--colspan用来说明跨列数-->
</tr>
<tr>
  <td>Bill Gates</td>
  <td>555 77 854</td>
  <td>555 77 855</td>
</tr>
</table>

<h4>单元格跨两行:</h4>
<table border="1">
<tr>
  <th>First Name:</th>
  <td>Bill Gates</td>
</tr>
<tr>
  <th rowspan="2">Telephone:</th><!--rowspan用来说明跨行数-->
  <td>555 77 854</td>
</tr>
<tr>
  <td>555 77 855</td>
</tr>
</table>

</body>
</html>
```
### 表格的嵌套
```html
<!DOCTYPE html>
<html>
<head> 
<meta charset="utf-8"> 
<title>AHUTSJ_HTML</title> 
</head>
<body>

<table border="1">
<tr>
  <td>
   <p>这是一个段落</p>
   <p>这是另一个段落</p>
  </td>
  <td>这个单元格包含一个表格:
   <table border="1">
   <tr>
     <td>A</td>
     <td>B</td>
   </tr>
   <tr>
     <td>C</td>
     <td>D</td>
   </tr>
   </table>
  </td>
</tr>
<tr>
  <td>这个单元格包含一个列表
   <ul>
    <li>apples</li>
    <li>bananas</li>
    <li>pineapples</li>
   </ul>
  </td>
  <td>HELLO</td>
</tr>
</table>

</body>
</html>
```
## HTML列表
HTML 支持有序、无序和定义列表。
### 无序列表
无序列表是一个项目的列表，此列项目使用粗体圆点（典型的小黑圆圈）进行标记。无序列表使用 ul 标签。
```html
<ul>
<li>Coffee</li>
<li>Milk</li>
</ul>
```
### 有序列表
有序列表也是一列项目，列表项目使用数字进行标记。 有序列表始于 ol 标签。每个列表项始于 li 标签,列表项使用数字来标记。
```html
<ol>
<li>Coffee</li>
<li>Milk</li>
</ol>
```
### 自定义列表
自定义列表不仅仅是一列项目，而是项目及其注释的组合。
自定义列表以 dl 标签开始，每个自定义列表项以 dt 开始，每个自定义列表项的定义以 dd 开始：
```html
<dl>
<dt>Coffee</dt>
<dd>- black hot drink</dd>
<dt>Milk</dt>
<dd>- white cold drink</dd>
</dl>
```
## HTML div 和 span
HTML可以通过 div 和 span 将元素组合起来。
### HTML区块元素
大多苏HTML元素被定义为块级元素或内联元素。
块级元素在浏览器显示时，通常会以新行来开始（和结束）。
实例： h1 p ul table
### HTML内联元素
内联元素在显示时通常不会以新行开始。
实例：b td a img
## HTML布局
网页布局对改善网站的外观非常重要。
### 使用 div 元素的网页布局
```html
<!DOCTYPE html>
<html>
<head> 
<meta charset="utf-8"> 
<title>AHUTSJ_HTML</title> 
</head>
<body>

<div id="container" style="width:500px">

<div id="header" style="background-color:#FFA500;">
<h1 style="margin-bottom:0;">主要的网页标题</h1></div>

<div id="menu" style="background-color:#FFD700;height:200px;width:100px;float:left;">
<b>菜单</b><br>
HTML<br>
CSS<br>
JavaScript</div>

<div id="content" style="background-color:#EEEEEE;height:200px;width:400px;float:left;">
内容在这里</div>

<div id="footer" style="background-color:#FFA500;clear:both;text-align:center;">
版权 © 1466291924@qq.com</div>

</div>
 
</body>
</html>
```
### 使用table元素的网页布局
```html
<!DOCTYPE html>
<html>
<head> 
<meta charset="utf-8"> 
<title>AHUTSJ_HTML</title> 
</head>
<body>

<table width="500" border="0">
<tr>
<td colspan="2" style="background-color:#FFA500;">
<h1>主要的网页标题</h1>
</td>
</tr>

<tr>
<td style="background-color:#FFD700;width:100px;vertical-align:top;">
<b>菜单</b><br>
HTML<br>
CSS<br>
JavaScript
</td>
<td style="background-color:#eeeeee;height:200px;width:400px;vertical-align:top;">
内容在这里</td>
</tr>

<tr>
<td colspan="2" style="background-color:#FFA500;text-align:center;">
版权 © 1466291924@qq.com</td>
</tr>
</table>

</body>
</html>
```
**提示**：虽然我们可以使用HTML table标签来设计出漂亮的布局，但是table标签是不建议作为布局工具使用的，表格不是布局工具。
### HTML布局的补充说明
1. 使用 CSS 最大的好处是，如果把 CSS 代码存放到外部样式表中，那么站点会更易于维护。通过编辑单一的文件，就可以改变所有页面的布局。
2. 由于创建高级的布局非常耗时，使用模板是一个快速的选项。通过搜索引擎可以找到很多免费的网站模板（您可以使用这些预先构建好的网站布局，并优化它们）
## HTML表单和输入
HTML表单用于收集不同类型的用户输入。表单元素是允许用户在表单中输入内容,比如：文本域(textarea)、下拉列表、单选框(radio-buttons)、复选框(checkboxes)等等。表单使用 form 标签来设置。
### HTML表单——文本域（text-fields）
文本域通过 input type=“text” 标签来设定，当用户要在表单中键入字母、数字等内容时，就会用到文本域。
```html
<form>
First name: <input type="text" name="firstname"><br>
Last name: <input type="text" name="lastname">
</form>
```
**注意**：表单本身并不可见。同时，在大多数浏览器中，文本域的默认宽度是 20 个字符。
### HTML表单——密码字段
密码字段通过标签 input type=“password” 来定义：
```html
<form>
Password: <input type="password" name="pwd">
</form>
```
**注意**：密码字段不会明文显示，而是以星号或者原点代替。
### HTML表单——单选按钮（radio buttons）
input type=“radio” 标签定义了表单单选框选项。
```html
<form>
<input type="radio" name="sex" value="male">Male<br>
<input type="radio" name="sex" value="female">Female
</form>
```
### HTML表单——复选框（checkboxes）
input type=“checkboxes” 定义了复选框，用户需要从若干给定的选择中选取一个或若干选项。
```html
<form>
<input type="checkbox" name="vehicle" value="Bike">I have a bike<br>
<input type="checkbox" name="vehicle" value="Car">I have a car
</form>
```
### HTML表单——提交按钮（submit button）
input type=“submit” 定义了提交按钮。当用户单击确认按钮时，表单的内容会被传送到另一个文件。表单的动作属性定义了目的文件的文件名。由动作属性定义的这个文件通常会对接收到的输入数据进行相关的处理。
```html
<form name="input" action="html_form_action.php" method="get">
Username: <input type="text" name="user">
<input type="submit" value="Submit">
</form>
```
假如您在上面的文本框内键入几个字母，然后点击确认按钮，那么输入数据会传送到 "html_form_action.php" 的页面。该页面将显示出输入的结果。
### HTML表单——下拉菜单（带预选）
```html
<!DOCTYPE html>
<html>
<head> 
<meta charset="utf-8"> 
<title>AHUTSJ_HTML</title> 
</head>
<body>

<form action="">
<select name="cars">
<option value="volvo">Volvo</option>
<option value="saab">Saab</option>
<option value="fiat" selected>Fiat</option><!--selected表示预选项，去掉则无预选，默认显示第一项-->
<option value="audi">Audi</option>
</select>
</form>

</body>
</html>
```
### 文本域
```html
<!DOCTYPE html>
<html>
<head> 
<meta charset="utf-8"> 
<title>AHUTSJ_HTML</title> 
</head>
<body>

<textarea rows="10" cols="30">
我是一个文本框。
</textarea>

</body>
</html>
```
## HTML框架
通过使用框架，可以在一个浏览器窗口显示不止一个界面。
框架的定义格式：
`<iframe src="URL"></iframe>`
该URL指向不同的网页。
### Iframe - 设置高度与宽度
height 和 width 属性用来定义iframe标签的高度与宽度。属性默认以像素为单位, 但是你可以指定其按比例显示 (如："80%")。
`<iframe src="demo_iframe.htm" width="200" height="200"></iframe>`
### Iframe - 移除边框
设置属性值为 "0" 移除iframe的边框:
`<iframe src="demo_iframe.htm" frameborder="0"></iframe>`
## HTML颜色
颜色名和颜色值的内容 :#000000-#FFFFFF(RGB)
## HTML字符实体
有关HTML字符实体的内容，可在回忆不起来时查表得知。




