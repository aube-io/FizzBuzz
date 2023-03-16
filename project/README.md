# Documentation du projet

## Structure du projet
```
project
├── cmd 			Point d'entrée de l'application
├── config 			Configuration du projet
├── internal 			Packages internes
│   ├── api 			Packages réserver à la création d'API
│   │   └── v1 			Version des API
│   │       ├── fizzbuzz 	API FizzBuzz
│   │       └── status   	API Status
│   ├── docs			Documentation Swagger
│   ├── lib			Libraries
│   └── model 			Pseudo model de données
└── pkg 			Pour les packages exportables
```

## Approche
Comme je ne savais pas trop ce qui étais attendu j'ai essayer d'édulcorer le projet avec différentes approches et bonne pratiques
C'est pas toujours ce qu'il y aura de plus optimiser/performant mais ca fait le job.