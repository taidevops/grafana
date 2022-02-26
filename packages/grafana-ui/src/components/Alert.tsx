import React, { HTMLAttributes, ReactNode } from 'react';
import { css, cx } from '@emotion/css';

export type AlertVariant = 'success' | 'warning' | 'error' | 'info';

export interface Props extends HTMLAttributes<HTMLDivElement> {
  title: string;
}

export const Alert = React.forwardRef<HTMLDivElement, Props>(
  (
    {
      title
    },
    ref
  ) => {
    const styles = getStyles();
    return (
      <div
        ref={ref}
        className={cx(styles.alert)}
      >
        <div>
          <i></i>
        </div>
      </div>
    )
  }
)

const getStyles = (

) => {

  return {
    alert: css``,
    icon: css``,
    title: css``,
  }
}
