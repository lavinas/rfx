class Cadoc6334Discount:
    # valores brutos
    transaction_amount: float
    transaction_quantity: int
    # valores auxiliares para cálculo do desvio padrão
    sqrdiff_mcc_fee: float
    # valores estatísticos
    avg_mcc_fee: float
    min_mcc_fee: float
    max_mcc_fee: float
    stdev_mcc_fee: float

    # construtor
    def __init__(self):
        self.transaction_amount = 0.0
        self.transaction_quantity = 0
        self.avg_mcc_fee = 0.0
        self.min_mcc_fee = 0.0
        self.max_mcc_fee = 0.0
        self.sqrdiff_mcc_fee = 0.0
        self.stdev_mcc_fee = 0.0

    # atualiza os valores do desconto com base no valor da transação e valor do mcc
    def update_values(self, transaction_amount: float, transaction_mcc_value: float):
        # atualiza valor
        self.transaction_amount += transaction_amount
        # atualiza quantidade
        self.transaction_quantity += 1
        # calcula taxa baseado no valor da transação e o valor do mcc
        mcc_fee = round(transaction_mcc_value / transaction_amount * 100, 2)
        # se é a primeira transação, inicializa os valores
        if self.transaction_quantity == 1:
            self.min_mcc_fee = mcc_fee
            self.max_mcc_fee = mcc_fee
            self.avg_mcc_fee = mcc_fee
            self.sqrdiff_mcc_fee = 0.0
            self.stdev_mcc_fee = 0.0
            return
        # se não é a primeira transação, atualiza os valores
        # atualiza o valor mínimo
        self.min_mcc_fee = min(self.min_mcc_fee, mcc_fee)
        # atualiza o valor máximo
        self.max_mcc_fee = max(self.max_mcc_fee, mcc_fee)
        # atualiza a média e desvio padrão pelo algoritmo de Welford
        # calcula a media
        delta = mcc_fee - self.avg_mcc_fee
        self.avg_mcc_fee += delta / self.transaction_quantity
        # arredonda a média para 2 casas decimais
        self.avg_mcc_fee = round(self.avg_mcc_fee, 2)
        # calcular diferenca quadratica
        delta2 = mcc_fee - self.avg_mcc_fee
        self.sqrdiff_mcc_fee += delta * delta2
        # calcular a variancia - soma das diferencas quadráticas dividido pela quantidade de transações menos 1
        var = self.sqrdiff_mcc_fee / (self.transaction_quantity - 1)
        # calcula o desvio padrão - raiz da variancia
        self.stdev_mcc_fee = var ** 0.5
        # arredonda o desvio padrão para 2 casas decimais
        self.stdev_mcc_fee = round(self.stdev_mcc_fee, 2)

# main - testar valores
if __name__ == "__main__":
    # base de teste
    test_values = [
        (100.0, 2.0), # 2%
        (200.0, 4.0), # 2%
        (7.0, 3.0), # 2%
        (100.0, 2.0), # 2%
        (1000.0, 1.5),  # 3%
        (500.0, 10.0), # 2%
        (50.0, 1.0), # 2%
        (300.0, 6.0), # 2%
        (150.0, 3.0), # 2%
        (80.0, 1.6) # 2%
    ]
    discount = Cadoc6334Discount()
    for amount, mcc in test_values:
        discount.update_values(amount, mcc)
        print(f"Transaction amount: {amount}, MCC value: {mcc}")
        print(f"Result: Amount {discount.transaction_amount}, \
              Quantity: {discount.transaction_quantity}, \
              Avg MCC Fee: {discount.avg_mcc_fee}%, \
              Min MCC Fee: {discount.min_mcc_fee}%, \
              Max MCC Fee: {discount.max_mcc_fee}%, \
              Stdev MCC Fee: {discount.stdev_mcc_fee}%\n")
 

