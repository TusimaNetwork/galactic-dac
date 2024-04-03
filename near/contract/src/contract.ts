// Find all our documentation at https://docs.near.org
import { NearBindgen, initialize, near, call, view, LookupMap } from 'near-sdk-js';

@NearBindgen({})
class NearStore {
  counting: bigint = BigInt(0);
  daStore: LookupMap<any>;
  
  constructor() {
    this.daStore = new LookupMap("s");
  }
   
  @view({}) // This method is read-only and can be called for free
  get_greett({ greeting }: { greeting: string }): bigint {
    return this.daStore.get(greeting);
  }

  @call({ privateFunction: true })
  set_greett({ greeting, greetingValue }: { greeting: string, greetingValue: string }): void {
    near.log(`Saving greeting ${greeting}`);
    // if(greetingValue.length > 16000){
    //   let count = greetingValue.length / 16000;
    //   for (let i=0; i<count; i++){
    //     let start = i*16000;
    //     let end = (i+1)*16000;
    //     if (end > greetingValue.length) {
    //       end = greetingValue.length;
    //     }
    //     let tmp = greetingValue.slice(start, end);
    //     near.log(`${tmp.length}`);
    //     // this.set_greett2({"greetingValue":tmp})
    //   }
    // }else{
    //   near.log(`${greetingValue}`);
    // }
    if(greetingValue.length <= 16000) {
      near.log(`${greetingValue}`);
    }
    this.counting ++
    this.daStore.set(greeting, near.blockHeight())
  }
}
