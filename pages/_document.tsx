import Document, { Html, Head, Main, NextScript } from 'next/document';
import React from 'react';

import Favicon from '../components/Favicon';
import MetaTag from '../components/MetaTag';
import config from '../ssg.config';

class MyDocument extends Document {
  render(): JSX.Element {
    return (
      <Html lang={config.lang}>
        <link
          rel="alternate"
          type="application/rss+xml"
          title={config.title}
          href={config.url + '/feed.xml'}
        />
        <Head>
          <Favicon />
          <MetaTag
            language={config.lang}
            description={config.description}
            properties={config.metaProperties}
          />
        </Head>
        <body>
          <Main />
          <NextScript />
        </body>
      </Html>
    );
  }
}

export default MyDocument;
