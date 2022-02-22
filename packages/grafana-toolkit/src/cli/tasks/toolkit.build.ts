import execa = require('execa');
import * as fs from 'fs';
import chalk from 'chalk';
import { useSpinner } from '../utils/useSpinner';
import { Task, TaskRunner } from './task';

const path = require('path');

let distDir: string, cwd: string;

const clean = () => useSpinner('Cleaning', () => execa('npm', ['run', 'clean']));

const compile = () =>
  useSpinner('Compiling sources', async () => {
    try {
      await execa('tsc', ['-p', './tsconfig.json']);
    } catch (e) {
      console.log(e);
      throw e;
    }
  });

  const savePackage = ({ path, pkg }: { path: string; pkg: {} }) =>
  useSpinner('Updating package.json', async () => {
    new Promise<void>((resolve, reject) => {
      fs.writeFile(path, JSON.stringify(pkg, null, 2), (err) => {
        if (err) {
          reject(err);
          return;
        }
        resolve();
      });
    });
  });

const preparePackage = async (pkg: any) => {
  pkg.bin = {
    'grafana-toolkit': './bin/grafana-toolkit.js',
  };

  await savePackage({
    path: `${cwd}/dist/package.json`,
    pkg,
  });
};

const copyFiles = () => {
  const files = [
    'bin/grafana-toolkit.js',
  ];

  return useSpinner(`Moving ${files.join(', ')} files`, async () => {
    const promises = files.map((file) => {
      return new Promise<void>((resolve, reject) => {
        const basedir = path.dirname(`${distDir}/${file}`);
        if (!fs.existsSync(basedir)) {
          fs.mkdirSync(basedir, { recursive: true });
        }
        fs.copyFile(`${cwd}/${file}`, `${distDir}/${file}`, (err) => {
          if (err) {
            reject(err);
            return;
          }
          resolve();
        });
      });
    });

    await Promise.all(promises);
  });
};

interface ToolkitBuildOptions {}

const toolkitBuildTaskRunner: TaskRunner<ToolkitBuildOptions> = async () => {
  cwd = path.resolve(__dirname, '../../../');
  distDir = `${cwd}/dist`;
  const pkg = require(`${cwd}/package.json`);
  console.log(chalk.yellow());

  await clean();
  await compile();
  await preparePackage(pkg);
  fs.mkdirSync('./dist/bin');
  fs.mkdirSync('./dist/sass');
  await copyFiles();
}

export const toolkitBuildTask = new Task<ToolkitBuildOptions>('@grafana/toolkit build', toolkitBuildTaskRunner);