// Find all our documentation at https://docs.near.org
import { NearBindgen, initialize, near, call, view, LookupMap } from 'near-sdk-js';

@NearBindgen({})
class HelloNear2 {
  counting: bigint = BigInt(0);
  daStore: LookupMap<any>;
  
  constructor() {
    this.daStore = new LookupMap("s");
  }
   
  @view({}) // This method is read-only and can be called for free
  get_greett({ greeting }: { greeting: string }): bigint {
    return this.daStore.get(greeting);
  }

  @call({}) // This method changes the state, for which it cost gas
  set_greett({ greeting, greetingValue }: { greeting: string, greetingValue: string }): void {
    near.log(`Saving greeting ${greeting}`);
    near.log(`${greetingValue}`);

    this.counting ++
    this.daStore.set(greeting, near.blockHeight())
  }

  @call({}) // This method changes the state, for which it cost gas
  set_greett2({ greeting }: { greeting: string }): void {
    near.log(`Saving greeting ${greeting}`);

    this.counting ++
    this.daStore.set(greeting, near.blockHeight())
  }
}
