import Head from 'next/head';
import React, { FC } from 'react';

import Layout from '../../components/Layout';
import Date from '../../components/MyDate';
import {
  getAllPostSlugs,
  getPostData,
  params,
  postData,
} from '../../lib/posts';
import config from '../../ssg.config';
import utilStyles from '../../styles/utils.module.css';

import styles from './slug.module.css';

type Props = {
  postData: postData;
};

const Post: FC<Props> = ({ postData }) => {
  return (
    <Layout home={false}>
      <>
        <Head>
          <title>
            {postData.title} | {config.title}
          </title>
          <meta name="twitter:description" content={postData.title} />
        </Head>
        <article className={styles.article}>
          <h1 className={utilStyles.headingXl}>{postData.title}</h1>
          <Date raw={postData.date} />
          <div dangerouslySetInnerHTML={{ __html: postData.contentHtml }} />
        </article>
      </>
    </Layout>
  );
};

export default Post;

export const getStaticPaths = async (): Promise<{
  paths: { params: params }[];
  fallback: boolean;
}> => {
  const paths = await getAllPostSlugs();
  return {
    paths,
    fallback: false,
  };
};

export const getStaticProps = async ({
  params,
}: {
  params: params;
}): Promise<{ props: { postData: { slug: string; contentHtml: string } } }> => {
  const postData = await getPostData(params.slug);
  return {
    props: {
      postData,
    },
  };
};
