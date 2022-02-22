import { Unsubscribable, Observable } from 'rxjs';

/**
 * @alpha
 * internal interface
 */
export interface BusEvent {
  readonly type: string;
  readonly payload?: any;
  readonly origin?: EventBus;
}

/**
 * @alpha
 * Main minimal interface
 */
 export interface EventBus {
  /**
   * Publish single vent
   */
  publish<T extends BusEvent>(event: T): void;
}