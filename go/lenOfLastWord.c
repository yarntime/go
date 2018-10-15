int lengthOfLastWord(char* s) {
    int len = strlen(s);
    if (len == 0) {
        return 0;
    }
    int result = 0;
    int pos = len - 1;
    while (s[pos] == ' ') pos--;
    while (pos >= 0 && s[pos] != ' ') {
        result++;
        pos--;
    } 
    return result;
}