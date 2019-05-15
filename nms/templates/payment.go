package templates

var TemplatePayment = `
	<h1>{{ .username }}</h1>
	<p>Seu pagamento no valor de {{ .transferValue }}foi realizado</p>
	<h3>Favorecido</h3>
	<p>{{ .benefited }}</p>
	<h3>Vencimento</h3>
	<p>{{ .time }}</p>
`
