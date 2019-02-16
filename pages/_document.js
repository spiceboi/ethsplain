// _document is only rendered on the server side and not on the client side
// Event handlers like onClick can't be added to this file

// ./pages/_document.js
import Document, { Head, Main, NextScript } from 'next/document'

class MyDocument extends Document {
  static async getInitialProps (ctx) {
    const initialProps = await Document.getInitialProps(ctx)
    return { ...initialProps }
  }

  render () {
    return (
      <html lang='en'>
        <Head>
          <meta charset='utf-8' />
          <meta name='viewport' content='width=device-width, initial-scale=1, shrink-to-fit=no' />
          <link rel='stylesheet' href='//maxcdn.bootstrapcdn.com/bootswatch/2.3.1/cyborg/bootstrap.min.css' />
          <link rel='stylesheet' href='https://use.fontawesome.com/releases/v5.7.2/css/all.css' integrity='sha384-fnmOCqbTlWIlj8LyTjo7mOUStjsKC4pOpQbqyi7RrhN7udi9RwhKkMHpvLbHG9Sr' crossorigin='anonymous' />
          <script src='//cdnjs.cloudflare.com/ajax/libs/jquery/3.3.1/jquery.min.js' />
          <script src='//cdnjs.cloudflare.com/ajax/libs/underscore.js/1.9.1/underscore-min.js' />
          <script src='//cdnjs.cloudflare.com/ajax/libs/d3/3.1.6/d3.min.js' />
          <script src='/stackpath.bootstrapcdn.com/bootstrap/4.3.1/js/bootstrap.min.js' integrity='sha384-JjSmVgyd0p3pXB1rRibZUAYoIIy6OrQ6VrjIEaFf/nJGzIxFDsf4x0xIM+B07jRM' crossorigin='anonymous' />
        </Head>
        <body data-theme='dark'>
          <Main />
          <NextScript />
          <script src='/static/ethsplain.js' />
        </body>
      </html>
    )
  }
}

export default MyDocument
