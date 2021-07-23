import React, { FC } from "react";

type Props = {
  language: string;
  description: string;
  properties?: {
    property: string;
    content: string;
  }[];
};

const MetaTag: FC<Props> = ({ language, description, properties = [] }) => (
  <>
    <meta name="viewport" content="initial-scale=1.0, width=device-width" />
    <meta charSet="UTF-8" />
    <meta httpEquiv="X-UA-Compatible" content="IE=edge" />
    <meta httpEquiv="content-language" content={language} />
    <meta
      name="viewport"
      content="width=device-width, initial-scale=1.0, maximum-scale=1.0, minimum-scale=1.0"
    />
    <meta name="description" content={description} />
    {properties.map((p) => (
      <meta key={p.property} property={p.property} content={p.content} />
    ))}
  </>
);

export default MetaTag;
