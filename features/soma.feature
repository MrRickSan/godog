#language: pt

Funcionalidade: Soma

    # Cenario: Soma de valores

    #     Quando faco a soma de 2 + 2
    #     Então o resultado deve ser 4

    Esquema do Cenario: Soma de valores

        Quando faco a soma de <v1> + <v2>
        Então o resultado deve ser <resultado>

        Exemplos:
        | v1 | v2 | resultado |
        | 2  | 2  | 4         |
        | 3  | 3  | 6         |
        | 2  | 5  | 7         |