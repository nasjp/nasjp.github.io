import { AppProps } from "next/app";
import React, { FC } from "react";

import "../styles/globals.css";
import "highlight.js/styles/googlecode.css";

const MyApp: FC<AppProps> = ({ Component, pageProps }) => (
  <>
    <Component {...pageProps} />
  </>
);

export default MyApp;
