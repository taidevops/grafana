import React from 'react';
import { GrafanaApp } from './app';

interface AppWrapperProps {
  app: GrafanaApp;
}


export class AppWrapper extends React.Component<AppWrapperProps> {
  constructor(props: AppWrapperProps) {
    super(props);
  }
  render() {
    return (<></>)
  }
}