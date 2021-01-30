## Revision

Revision will receive a .md file (or txt file) and will search for a word/phrase marked
as bold (using the "keyword" markdown uses to make words bold, \*\*) and group all together
and save it in another file. Then you can lookup the things you considered most important.

For example, if passed this file:
```markdown

## Example of markdown

Just **call on** me and I'll send it along
With love, from me to you
I've got everything that **you want**
Like a heart that's oh so true
Just call on me and **I'll send it along**
**With love**, from me to you

* Paul
* George
* Ringo
* **John**
```

will generate the following file

```txt
call on
you want
I'll send it along
With love
John
```

## TODO

Ter a possibilidade (talvez uma flag? -c (code)) de que o revision, alem do bold tambem considere textos
marcados com  o \`.
