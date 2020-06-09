## Que doit faire notre table

- input directement depuis json => traiter les NaN, les json dates, etc.
- input directement depuis un csv
- faire des calculs (expr)
- formatter la table avec les options: précision, numeral, etc.
- types de colonnes : 
  - int
  - uint
  - bool
  - float
  - decimal
  - string
  - time
  - serie....

  [ N°commande, Produit, Prix ]
  A, toto, 10
  A, tata, 15
  A, titi , 347
  B, lionel, 3568


[N° commande, Produits]
A, [{nom, prix}, {nom, prix}]

=> flatten(table)


A, toto, 10
A, tata, 15
A, titi , 347
B, lionel, 3568
