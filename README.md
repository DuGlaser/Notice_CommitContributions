# Notice_CommitContributions

## GitHubからCommit数を取得
```
query {
  user(login: "DuGlaser") {
    contributionsCollection(from: "2019-12-02T00:00:00", to: "2019-12-03T00:00:00") {
      totalCommitContributions # コミットした全レポジトリに対する総コミット数
    }
  }
}
```
