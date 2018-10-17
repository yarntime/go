const LETTER_COUNT = 26

type Trie struct {
	IsWord bool
	Next [LETTER_COUNT]*Trie
}


/** Initialize your data structure here. */
func Constructor() Trie {
	return Trie{
		IsWord: false,
		Next: [LETTER_COUNT]*Trie{},
	}
}


/** Inserts a word into the trie. */
func (this *Trie) Insert(word string)  {
	for i := 0; i < len(word); i++ {
		if this.Next[word[i] - 'a'] == nil {
			this.Next[word[i] - 'a'] = &Trie{
				IsWord: false,
				Next: [LETTER_COUNT]*Trie{},
			}
		}
		this = this.Next[word[i] - 'a']
	}
	this.IsWord = true
}


/** Returns if the word is in the trie. */
func (this *Trie) Search(word string) bool {
	for i := 0; i < len(word); i++ {
		if this.Next[word[i] - 'a'] == nil {
			return false
		}
		this = this.Next[word[i] - 'a']
	}
	return this.IsWord
}


/** Returns if there is any word in the trie that starts with the given prefix. */
func (this *Trie) StartsWith(prefix string) bool {
	for i := 0; i < len(prefix); i++ {
		if this.Next[prefix[i] - 'a'] == nil {
			return false
		}
		this = this.Next[prefix[i] - 'a']
	}
	return true
}