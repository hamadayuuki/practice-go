# react-todo

## セットアップ

VSCode 拡張機能インストール : ES7+React/Redux~, prettier

**フォルダ新規作成**

```.env
REACT_APP_API_URL=http://localhost:8080
```

```.prettierrc
{
    "singleQuote": true,
    "semi": false
}
```


**ターミナル**

```.bash
// Reactのライブラリ
$ npm i @tanstack/react-query@4.28.0
$ npm i @tanstack/react-query-devtools@4.28.0
$ npm i zustand@4.3.6
$ npm i @heroicons/react@2.0.16
$ npm i react-router-dom@6.10.0 axios@1.3.4

// tailwindcssインストール
// https://tailwindcss.com/docs/guides/create-react-app
$ npm install -D tailwindcss
$ npx tailwindcss init
```


**tailwindcss を有効にする**
```App.css
@tailwind base;
@tailwind components;
@tailwind utilities;
```


**npm start**
```
$ npm start
```
