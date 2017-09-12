"""
    levenshtein distance
    ~~~~~~~~~~~~~~~~~~~
    calculate levenshtein distance between two strings.
    :copyright: (c) 2017 by Blurt Heart.
    :license: BSD, see LICENSE for more details.
"""

def steps_to_convert(line1, line2):
    # Levenshtein Distance Algorithm
    length1 = len(line1)
    length2 = len(line2)
    # create two-dimensional array(length1+1, length2+1)
    first_row = [i for i in range(length2+1)]
    other_rows = [([i] * (length2+1)) for i in range(1, length1+1)]
    steps = [first_row] + other_rows

    for i in range(1, length1+1):
        for j in range(1, length2+1):
            min_step = min(steps[i-1][j-1], steps[i][j-1], steps[i-1][j])
            steps[i][j] = min_step if line1[i-1] == line2[j-1] else min_step + 1
    return steps[-1][-1]


if __name__ == "__main__":
    #These "asserts" using only for self-checking and not necessary for auto-testing
    assert steps_to_convert('line1', 'line1') == 0, "eq"
    assert steps_to_convert('line1', 'line2') == 1, "2"
    assert steps_to_convert('line', 'line2') == 1, "none to 2"
    assert steps_to_convert('ine', 'line2') == 2, "need two more"
    assert steps_to_convert('line1', '1enil') == 4, "everything is opposite"
    assert steps_to_convert('', '') == 0, "two empty"
    assert steps_to_convert('l', '') == 1, "one side"
    assert steps_to_convert('', 'l') == 1, "another side"