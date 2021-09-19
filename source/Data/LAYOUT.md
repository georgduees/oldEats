# Server Layout
# Source Data Layout
/data -- contains files
public folder:
/data/books/ID/01 name/[01-255]/image.jpeg|content.txt
# API-Layout
/books/content

# Inhaltsverzeichnis in Tabelle und Referenzen zur Seite.
tblBook
    ID,
    name,
    author,
    year,
    isbn
tblTOC
    ID,
    bookID,
    key,[Number]
    value,
    pageID
tblPage
    ID,
tblContent
    ID,text,
tblPosition

tblType
    ID,
    name
tblRecipe
    ID,
    name,

tblCategory
    ID,
    name,
