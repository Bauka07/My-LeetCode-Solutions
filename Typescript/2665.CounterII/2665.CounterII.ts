type Counter = {
  increment: () => number;
  decrement: () => number;
  reset: () => number;
};

function createCounterII(init: number): Counter {
  const start = init;
  return {
    increment: () => ++init,
    decrement: () => --init,
    reset: () => (init = start),
  };
}
