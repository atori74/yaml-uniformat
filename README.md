# yaml-uniformat

yamlを一意の形式にフォーマットし出力する

prettifyとの違い
yamlに含まれる配列を再帰的に文字列の昇順でソートする
->配列の順序は意味を持たないものとしてソートする

## 実行例
```sh
./yaml-uniformat < input.yaml
# output formatted yaml to stdout.
```

## Input
```input.yaml
people:
    - name: John
      age: 23
    - name: Robert
      age: 30
    - name: Alice
      age: 22
      _isLeader: true
    - "Tom"
    - 1
    - false
```

## Output
```output.yaml
people:
    - 1
    - Tom
    - _isLeader: true
      age: 22
      name: Alice
    - age: 23
      name: John
    - age: 30
      name: Robert
    - false
```

