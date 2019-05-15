package templates

var TemplateTransferUser = `
	<h1>Olá {{ .username}}</h1>
	<p>Você recebeu uma transferência de {{ .userFrom }}</p>
	<h3>Valor: R${{ .transferValue }}</h3>
	<p>{{.time}}</p>
`
