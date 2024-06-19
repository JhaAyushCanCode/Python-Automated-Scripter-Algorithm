def are_synonyms(word1, word2):
    synonyms = {
        "happy": ["joyful", "content", "cheerful"],
        "sad": ["unhappy", "sorrowful", "dejected"],
        "fast": ["quick", "swift", "rapid"],
        "slow": ["lethargic", "sluggish", "unhurried"],
        "big": ["large", "huge", "enormous"],
        "small": ["tiny", "little", "petite"],
        "smart": ["intelligent", "bright", "clever"],
        "dull": ["boring", "uninteresting", "tedious"]
    }
    
    if word1 in synonyms and word2 in synonyms[word1]:
        return True
    if word2 in synonyms and word1 in synonyms[word2]:
        return True
    return False

def main():
    word1 = input("Enter the first word: ").strip().lower()
    word2 = input("Enter the second word: ").strip().lower()
    
    if are_synonyms(word1, word2):
        print(f"'{word1}' and '{word2}' are synonyms!")
    else:
        print(f"'{word1}' and '{word2}' are not synonyms.")

if __name__ == "__main__":
    main()
