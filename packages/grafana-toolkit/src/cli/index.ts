// @ts-ignore
import chalk from 'chalk';
import { program } from 'commander';
import { execTask } from './utils/execTask';
import { toolkitBuildTask } from './tasks/toolkit.build';

export const run = (includeInternalScripts = false) => {
  if (includeInternalScripts) {
    program
      .command('toolkit:build')
      .description('Prepares grafana/toolkit dist package')
      .action(async (cmd) => {
        await execTask(toolkitBuildTask)({});
      });
  }

  program.option('-v, --version', 'Toolkit version').action(async () => {
    const version = '1';
    console.log(`v${version}`);
  });
}