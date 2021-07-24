import fs from 'fs';
import path from 'path';

import matter from 'gray-matter';
import remark from 'remark';
import highlight from 'remark-highlight.js';
import html from 'remark-html';

const postsDirectory = path.join(process.cwd(), 'posts');
const isDev = process.env.NODE_ENV === 'development';

export type matterData = {
  title: string;
  date: string;
  draft: boolean;
};

export type params = {
  slug: string;
};

export type postData = matterData &
  params & {
    contentHtml: string;
  };

export const getAllPostSlugs = async (): Promise<{ params: params }[]> => {
  const fileNames = fs.readdirSync(postsDirectory);

  const postDatas = await getPostDatas(fileNames);

  return postDatas.map((data) => {
    return { params: data };
  });
};

export const getAllSortedPostDatas = async (): Promise<postData[]> => {
  const fileNames = fs.readdirSync(postsDirectory);

  const postDatas = await getPostDatas(fileNames);

  return postDatas.sort((a, b) => {
    if (a.date < b.date) {
      return 1;
    } else {
      return -1;
    }
  });
};

export const getPostData = async (slug: string): Promise<postData> => {
  const fullPath = path.join(postsDirectory, `${slug}.md`);
  const fileContents = fs.readFileSync(fullPath, 'utf8');

  const matterResult = matter(fileContents);

  const processedContent = await remark()
    .use(html)
    .use(highlight)
    .process(matterResult.content);
  const contentHtml = processedContent.toString();

  const data = matterResult.data as matterData;

  return {
    slug,
    contentHtml,
    ...data,
  };
};

const getPostDatas = async (fileNames: string[]): Promise<postData[]> => {
  const postDatas = await Promise.all(
    fileNames.map(async (fileName) => {
      return await getPostData(fileName.split('.').slice(0, -1).join('.'));
    })
  );

  return postDatas.filter((postData) => isDev || !postData.draft);
};
