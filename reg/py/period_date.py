from datetime import date

# PARAMETROS
# tabela de fechamento de datas (banco de dados)
closing_date_table = {
    '2025/10': date(2025, 11, 2),
    '2025/11': date(2025, 12, 3),
    '2025/12': date(2025, 12, 28),
}
# dia default se não encontrar na tabela (configuração)
default_closing_day = 3

# ALGORITMO PRINICIPAL
# funcao exemplo que traz a data referente ao periodo da transacao 
def get_period_date(transaction_date: date, current_date: date, closing_date_table: dict) -> date:
    return_period_date = None
    # Extrai o ano e o mês da data da transação
    year = transaction_date.year
    month = transaction_date.month
    # busca na tabela de fechamento
    key = f"{year}/{month:02d}"
    closing_date = closing_date_table.get(key)
    # se não encontrar na tabela, pega o dia default que é o dia default com o mês/ano seguinte da transação
    if closing_date is None:
        closing_date = date(year + (month // 12), (month % 12) + 1, default_closing_day)
    # verifica que a data atual é menor à data de fechamento e retorna a data da transação
    if current_date < closing_date:
       return transaction_date
    # caso contrário, retorna a maior data entra a data atual e o primeiro dia do mês seguinte
    # isto para o caso da data de fechamento retroativa (dentro do mesmo mês), garantindo que a data do período seja do mes seguinte
    next_month_first_day = date(year + (month // 12), (month % 12) + 1, 1)
    return max(current_date, next_month_first_day)

# APOIO
# procedure apenas para imprimir o resultado
def print_result(transaction_date: date, current_date: date, period_date: date):
    print(f"Data da transação: {transaction_date}")
    print(f"Data atual: {current_date}")
    print(f"Data do período: {period_date}\n")

# TESTES
# main
if __name__ == "__main__":
    # teste 1: data atual antes da data de fechamento
    transaction_date = date(2025, 10, 20)
    current_date = date(2025, 11, 1)
    period_date = get_period_date(transaction_date, current_date, closing_date_table)
    print(f"Teste 1: data atual antes da data de fechamento:")
    print_result(transaction_date, current_date, period_date)
    # teste 2: data atual depois da data de fechamento
    transaction_date = date(2025, 10, 20)
    current_date = date(2025, 11, 5)
    period_date = get_period_date(transaction_date, current_date, closing_date_table)
    print(f"Teste 2: data atual depois da data de fechamento:")
    print_result(transaction_date, current_date, period_date)
    # teste 3: data de fechamento não encontrada na tabela
    # data atual no mesmo dia do fechamento
    transaction_date = date(2025, 10, 20)
    current_date = date(2025, 11, 3)
    period_date = get_period_date(transaction_date, current_date, closing_date_table)
    print(f"Teste 3: data de fechamento não encontrada na tabela:")
    print_result(transaction_date, current_date, period_date)
    # data de fechamento retroativa e data atual antes da data de fechamento
    transaction_date = date(2025, 12, 15)
    current_date = date(2025, 12, 27)
    period_date = get_period_date(transaction_date, current_date, closing_date_table)
    print(f"Teste 4: data de fechamento retroativa e data atual antes da data de fechamento:")
    print_result(transaction_date, current_date, period_date)
    # data de fechamento retroativa e data atual depois da data de fechamento mas dentro do proprio mes
    transaction_date = date(2025, 12, 15)
    current_date = date(2025, 12, 29)
    period_date = get_period_date(transaction_date, current_date, closing_date_table)
    print(f"Teste 5: data de fechamento retroativa e data atual depois da data de fechamento mas dentro do proprio mes:")
    print_result(transaction_date, current_date, period_date)
    # data de fechamento retroativa e data atual depois da data de fechamento e no mes seguinte
    transaction_date = date(2025, 12, 15)
    current_date = date(2026, 1, 5)
    period_date = get_period_date(transaction_date, current_date, closing_date_table)
    print(f"Teste 6: data de fechamento retroativa e data atual depois da data de fechamento e no mes seguinte:")
    print_result(transaction_date, current_date, period_date)
    # data de fechamento não encontrada na tabela, data atual antes da data de fechamento
    transaction_date = date(2026, 1, 10)
    current_date = date(2026, 2, 2)
    period_date = get_period_date(transaction_date, current_date, closing_date_table)
    print(f"Teste 7: data de fechamento não encontrada na tabela, data atual antes da data de fechamento:")
    print_result(transaction_date, current_date, period_date)
    # data de fechamento não encontrada na tabela, data atual depois da data de fechamento
    transaction_date = date(2026, 1, 10)
    current_date = date(2026, 2, 5)
    period_date = get_period_date(transaction_date, current_date, closing_date_table)
    print(f"Teste 8: data de fechamento não encontrada na tabela, data atual depois da data de fechamento:")
    print_result(transaction_date, current_date, period_date)
    # data de fechamento não encontrada na tabela, data atual exatamente no dia de fechamento
    transaction_date = date(2026, 1, 10)
    current_date = date(2026, 2, 3)
    period_date = get_period_date(transaction_date, current_date, closing_date_table)
    print(f"Teste 9: data de fechamento não encontrada na tabela, data atual exatamente no dia de fechamento:")
    print_result(transaction_date, current_date, period_date)