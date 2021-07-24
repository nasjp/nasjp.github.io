import Head from 'next/head';
import Link from 'next/link';
import React, { FC } from 'react';

import Layout from '../components/Layout';
import Date from '../components/MyDate';
import { getAllSortedPostDatas, postData } from '../lib/posts';
import config from '../ssg.config';
import utilStyles from '../styles/utils.module.css';

export const getStaticProps = async (): Promise<{
  props: { postDatas: postData[] };
}> => {
  const postDatas = await getAllSortedPostDatas();
  return {
    props: {
      postDatas,
    },
  };
};

type Props = {
  postDatas: postData[];
};

const IndexPage: FC<Props> = ({ postDatas }) => {
  return (
    <Layout home={true}>
      <>
        <Head>
          <title>{config.title}</title>
          <meta name="twitter:description" content={config.title} />
        </Head>
        <section className={utilStyles.headingMd}></section>
        <section className={`${utilStyles.headingMd} ${utilStyles.padding1px}`}>
          <ul className={utilStyles.list}>
            {postDatas.map(({ slug, date, title }) => (
              <li className={utilStyles.listItem} key={slug}>
                <Link href={`/posts/${slug}`}>
                  <a className={utilStyles.colorInherit}>{title}</a>
                </Link>
                <small>
                  <Date raw={date} />
                </small>
              </li>
            ))}
          </ul>
        </section>
      </>
    </Layout>
  );
};

export default IndexPage;
