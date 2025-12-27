function twoSum(nums: number[], target: number): number[] | null {
  const result = new Map<number, number>();

  for (const [i, num] of nums.entries()) {
    const current = target - num;

    if (result.has(current)) {
      return [result.get(current)!, i];
    }

    result.set(num, i);
  }

  return null;
}
