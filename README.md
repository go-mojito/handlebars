<h1 align="center"><strong>Handlebars for Mojito</strong></h1>
<p align="center">
    <a href="https://goreportcard.com/report/github.com/go-mojito/handlebars" alt="Go Report Card">
        <img src="https://goreportcard.com/badge/github.com/go-mojito/handlebars" /></a>
	<a href="https://github.com/go-mojito/handlebars" alt="Go Version">
        <img src="https://img.shields.io/github/go-mod/go-version/go-mojito/handlebars.svg" /></a>
	<a href="https://godoc.org/github.com/go-mojito/handlebars" alt="GoDoc reference">
        <img src="https://img.shields.io/badge/godoc-reference-blue.svg"/></a>
	<a href="https://github.com/go-mojito/handlebars/blob/main/LICENSE" alt="Licence">
        <img src="https://img.shields.io/github/license/Ileriayo/markdown-badges?style=flat-square" /></a>
	<a href="https://makeapullrequest.com">
        <img src="https://img.shields.io/badge/PRs-welcome-brightgreen.svg?style=flat-square" alt="PRs Welcome"></a>
</p>
<p align="center">
    <a href="https://go.dev/" alt="Made with Go">
        <img src="https://ForTheBadge.com/images/badges/made-with-go.svg" /></a>
		
</p>
<p align="center">
Handlebars for Mojito provides a Handlebars Renderer implementation that was designed specifically for Mojito.
</p>

<p align="center"><strong>SonarCloud Report</strong></p>
<p align="center">
    <a href="https://sonarcloud.io/summary/overall?id=go-mojito_handlebars" alt="Quality Gate">
        <img src="https://sonarcloud.io/api/project_badges/measure?project=go-mojito_handlebars&metric=alert_status" /></a>
    <a href="https://sonarcloud.io/summary/overall?id=go-mojito_handlebars" alt="Quality Gate">
        <img src="https://sonarcloud.io/api/project_badges/measure?project=go-mojito_handlebars&metric=sqale_rating" /></a>
    <a href="https://sonarcloud.io/summary/overall?id=go-mojito_handlebars" alt="Quality Gate">
        <img src="https://sonarcloud.io/api/project_badges/measure?project=go-mojito_handlebars&metric=reliability_rating" /></a>
    <a href="https://sonarcloud.io/summary/overall?id=go-mojito_handlebars" alt="Quality Gate">
        <img src="https://sonarcloud.io/api/project_badges/measure?project=go-mojito_handlebars&metric=security_rating" /></a>
	<br>
    <a href="https://sonarcloud.io/summary/overall?id=go-mojito_handlebars" alt="Quality Gate">
        <img src="https://sonarcloud.io/api/project_badges/measure?project=go-mojito_handlebars&metric=vulnerabilities" /></a>
    <a href="https://sonarcloud.io/summary/overall?id=go-mojito_handlebars" alt="Quality Gate">
        <img src="https://sonarcloud.io/api/project_badges/measure?project=go-mojito_handlebars&metric=code_smells" /></a>
    <a href="https://sonarcloud.io/summary/overall?id=go-mojito_handlebars" alt="Quality Gate">
        <img src="https://sonarcloud.io/api/project_badges/measure?project=go-mojito_handlebars&metric=bugs" /></a>
</p>

<p align="center"><strong>How to register</strong></p>
<p align="center">To register the handlebars renderer as default, do the following in your main file</p>

```go
func init() {
    handlebars.AsDefault()
}
```
<p align="center">To register the handlebars renderer as a named renderer, do the following in your main file</p>

```go
func init() {
    handlebars.As("myRenderer")
}
```
