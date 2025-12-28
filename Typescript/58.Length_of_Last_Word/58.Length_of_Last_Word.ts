function lengthOfLastWord(s: string): number {
  if (s.length == 0) return 0;
  let l: number = 0;
  for (let i = s.length - 1; i >= 0; i--) {
    if (s[i] !== " ") l++;
    else if (l > 0) break;
  }
  return l;
}
