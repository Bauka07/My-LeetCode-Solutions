type ToBeOrNotToBe = {
  toBe: (val: any) => boolean;
  notToBe: (val: any) => boolean;
};

function expect(val: any): ToBeOrNotToBe {
  return {
    toBe: (num) => {
      if (val === num) return true;
      throw new Error("Not Equal");
    },
    notToBe: (num) => {
      if (val !== num) return true;
      throw new Error("Equal");
    },
  };
}
