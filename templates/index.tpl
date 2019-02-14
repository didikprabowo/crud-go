{{ define "index.html"}}
<!DOCTYPE html>
<html lang="en">

<head>
    <meta charset="UTF-8">
    <meta name="viewport" content="width=device-width, initial-scale=1.0">
    <meta http-equiv="X-UA-Compatible" content="ie=edge">
    <title>
        {{ .Title}}
    </title>

</head>

<body>
    {{ template "_header.html"}}
    {{if .Keyword -}}
    Your purchases:
    {{ range .Keyword }}
    {{ .}}
    {{ end}}
    {{else}}
    You didn't make any purchases during the period.
    {{end}}

    {{if .Ogp -}}
    Your purchases:
    {{ range .Ogp }}
    {{ .Url}}
    {{ end}}
    {{else}}
    You didn't make any purchases during the period.
    {{end}}
    {{ .Process}}
    {{range $index, $elem := .Ogp}}
    {{$elem }},
    {{end}}
</body>

</html>
{{ end}}