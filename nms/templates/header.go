package templates

var TemplateHeader = `
<!DOCTYPE html PUBLIC "-//W3C//DTD XHTML 1.0 Transitional//EN" "http://www.w3.org/TR/xhtml1/DTD/xhtml1-transitional.dtd">
<html xmlns="http://www.w3.org/1999/xhtml">
  <head>
    <meta name="viewport" content="width=device-width"/>
    <meta http-equiv="Content-Type" content="text/html; charset=UTF-8" />
    <title>nuCash</title>
    <style type="text/css">
      body{
        margin: 0 auto;
        padding: 0;
        min-width: 100%;
        font-family: sans-serif;
      }
      p{
        color: #424242;  
      }
      .content{
        max-width: 500px;
        margin: 0 auto;
      }
      header, footer{
        background: #424242;
        text-align: center;
        padding: 10px 0;
      }

      footer p {
        text-align: center;
        color: #fff;
      }
      
    </style>
  </head>
  <body>
    <div class="content">
      <header>
        <img class="logo" src="https://painel.nucash.com.br/img/nCs.png" alt="logo nucash">
      </header>`
