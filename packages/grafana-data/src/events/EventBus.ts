import EventEmitter from 'eventemitter3';
import { Unsubscribable, Observable } from 'rxjs';
import {
  EventBus,
  LegacyEmitter,
  BusEventHandler,
  BusEventType,
  LegacyEventHandler,
  BusEvent,
  AppEvent,
  EventFilterOptions,
} from './types';

export class EventBusSrv implements EventBus {
  private emitter: EventEmitter;

  constructor() {
    this.emitter = new EventEmitter();
  }

  publish<T extends BusEvent>(event: T): void {
    this.emitter.emit(event.type, event);
  }

  getStream<T extends BusEvent>(eventType: BusEventType<T>): Observable<T> {
    return new Observable<T>((observer) => {
      const handler = (event: T) => {
        observer.next(event);
      };

      this.emitter.on(eventType.type, handler);

      return () => {
        this.emitter.off(eventType.type, handler);
      };
    });
  }
}

class ScopedEventBus implements EventBus {
  
  constructor(public path: string[], private eventBus: EventBus) {
    
  }

  publish<T extends BusEvent>(event: T): void {
    if (!event.origin) {
      (event as any).origin = this;
    }
    this.eventBus.publish(event);
  }
}