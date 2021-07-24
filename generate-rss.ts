import fs from 'fs';
import { join } from 'path';

import { getAllSortedPostDatas, getPostData, postData } from './lib/posts';
import config from './ssg.config';

export const getPost = async (slug: string): Promise<postData> => {
  const postData = await getPostData(slug);
  return postData;
};

const createFeed = (post: postData) => `    <item>
      <title>${post.title}</title>
      <link>${config.url}/posts/${post.slug}</link>
      <guid>${config.url}/posts/${post.slug}</guid>
      <pubDate>${new Date(post.date).toUTCString()}</pubDate>
    </item>`;

const writeRss = async (filePath: string, content: string) => {
  try {
    await fs.promises.writeFile(filePath, content, 'utf-8');
  } catch (err) {
    console.error(err);
  }
};

const generateRss = async () => {
  const posts = await getAllSortedPostDatas();
  const lastBuildDate = new Date(posts.slice(-1)[0].date).toUTCString();
  const feeds = posts.map(post => createFeed(post));

  const rss = `<?xml version="1.0" ?>
<rss version="2.0" xmlns:atom="http://www.w3.org/2005/Atom">
  <channel>
    <title>${config.title}</title>
    <link>${config.url}</link>
    <description>${config.description}</description>
    <language>${config.lang}</language>
    <lastBuildDate>${lastBuildDate}</lastBuildDate>
    <atom:link href="${
      config.url
    }/feed.xml" rel="self" type="application/rss+xml" />
${feeds.join('\n')}
  </channel>
</rss>`;

  const filePath = join(process.cwd(), 'out/feed.xml');

  await writeRss(filePath, rss);
};

generateRss();
