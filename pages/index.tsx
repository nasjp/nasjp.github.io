import Head from 'next/head';
import Link from 'next/link';
import React, { FC } from 'react';

import Layout from '../components/Layout';
import { Date } from '../components/date';
import { params, matterData, getSortedPostsData } from '../lib/posts';
import config from '../ssg.config';
import utilStyles from '../styles/utils.module.css';

export const getStaticProps = async (): Promise<{
  props: { allPostsData: allPostsData };
}> => {
  const allPostsData = getSortedPostsData();
  return {
    props: {
      allPostsData,
    },
  };
};

type allPostsData = (matterData & params)[];

type Props = {
  allPostsData: allPostsData;
};

const IndexPage: FC<Props> = ({ allPostsData }) => {
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
            {allPostsData.map(({ slug, date, title }) => (
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
