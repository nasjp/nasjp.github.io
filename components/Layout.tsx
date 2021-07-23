import Link from 'next/link';
import React, { FC } from 'react';

import config from '../ssg.config';
import utilStyles from '../styles/utils.module.css';

import styles from './layout.module.css';

type Props = {
  children: JSX.Element;
  home: boolean;
};

const Layout: FC<Props> = ({ children, home = false }) => {
  return (
    <div className={styles.container}>
      <header className={styles.header}>
        {home ? (
          <>
            <h1 className={utilStyles.heading2Xl}>{config.title}</h1>
          </>
        ) : (
          <>
            <h1 className={utilStyles.heading2Xl}>
              <Link href="/">
                <a className={utilStyles.colorInherit}>{config.title}</a>
              </Link>
            </h1>
          </>
        )}
      </header>
      <main className={utilStyles.main}>{children}</main>
      <footer className={styles.footer}>&copy; {config.title}</footer>
    </div>
  );
};

export default Layout;
