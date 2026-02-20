class Cadoc6334Interchange:
    # valores brutos
    transaction_amount: float
    transaction_quantity: int
    # valores estatísticos
    avg_interchange_fee: float
 
    # construtor
    def __init__(self):
        self.transaction_amount = 0.0
        self.transaction_quantity = 0
        self.avg_interchange_fee = 0.0


    def update_values(self, transaction_amount: float, interchange_fee_value: float):
        # atualiza valor
        self.transaction_amount += transaction_amount
        # atualiza quantidade
        self.transaction_quantity += 1
        # calcula taxa baseado no valor da transação e o valor do mcc
        interchange_fee = round(interchange_fee_value / transaction_amount * 100, 2)
        # se é a primeira transação, inicializa os valores
        if self.transaction_quantity == 1:
            self.avg_interchange_fee = interchange_fee
            return
        # se não é a primeira transação, atualiza os valores
        # atualiza a média pelo algoritmo de Welford
        # calcula a media
        delta = interchange_fee - self.avg_interchange_fee
        self.avg_interchange_fee += delta / self.transaction_quantity
        # arredonda a média para 2 casas decimais
        self.avg_interchange_fee = round(self.avg_interchange_fee, 2)

# main
if __name__ == "__main__":
    # test value
    test_values = [
        (100.0, 1.0),
        (200.0, 3.5),
        (300.0, 7.8),
        (400.0, 9.6),
        (500.0, 5.0),
    ]
    interchange = Cadoc6334Interchange()
    for amount, fee in test_values:
        interchange.update_values(amount, fee)
        print(f"After transaction amount: {amount}, fee: {fee}")
        print(interchange.__dict__)
    