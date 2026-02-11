class Cadoc6334Discount:
    transaction_amount: float
    transaction_quantity: int
    avg_mcc_fee: float
    min_mcc_fee: float
    max_mcc_fee: float
    m2 : float
    stdev_mcc_fee: float

    # construtor
    def __init__(self):
        self.transaction_amount = 0.0
        self.transaction_quantity = 0
        self.avg_mcc_fee = 0.0
        self.min_mcc_fee = 0.0
        self.max_mcc_fee = 0.0
        self.m2 = 0.0
        self.stdev_mcc_fee = 0.0

    # atualiza os valores do desconto com base no valor e taxa da transação
    def update_values(self, transaction_amount: float, transaction_mcc_value: float):
        # atualiza valor
        self.transaction_amount += transaction_amount
        # atualiza quantidade
        self.transaction_quantity += 1
        # pega a taxa de desconto da transação
        mcc_fee = round(transaction_mcc_value / transaction_amount * 100, 2)
        # atualiza o valor mínimo
        self.min_mcc_fee = min(self.min_mcc_fee, mcc_fee) if self.transaction_quantity > 1 else mcc_fee
        # atualiza o valor máximo
        self.max_mcc_fee = max(self.max_mcc_fee, mcc_fee) if self.transaction_quantity > 1 else mcc_fee
        # atualiza a média e desvio padrão - algotimo de Welford
        delta = mcc_fee - self.avg_mcc_fee
        self.avg_mcc_fee += delta / self.transaction_quantity
        delta2 = mcc_fee - self.avg_mcc_fee
        self.m2 += delta * delta2
        self.stdev_mcc_fee = round((self.m2 / (self.transaction_quantity - 1)) ** 0.5, 2) if self.transaction_quantity > 1 else 0.0

# main - testar valores
if __name__ == "__main__":
    # base de teste
    test_values = [
        (100.0, 2.0), # 2%
        (200.0, 4.0), # 2%
        (150.0, 3.0), # 2%
        (100.0, 1.0), # 1%
        (50.0, 1.5),  # 3%
    ]
    discount = Cadoc6334Discount()
    for amount, mcc in test_values:
        discount.update_values(amount, mcc)
        print(f"Transaction amount: {amount}, MCC value: {mcc}")
        print(f"Result: {discount.transaction_amount}, \
              Quantity: {discount.transaction_quantity}, \
              Avg MCC Fee: {discount.avg_mcc_fee}%, \
              Min MCC Fee: {discount.min_mcc_fee}%, \
              Max MCC Fee: {discount.max_mcc_fee}%, \
              Stdev MCC Fee: {discount.stdev_mcc_fee}%\n")
 

