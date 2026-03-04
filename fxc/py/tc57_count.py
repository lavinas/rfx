from collections import Counter
import os

# Parametres
alvo = "../files/TC57"
# alvo = "../files/TC57/"
coluna_inicio = 49
coluna_fim = 68
lista = True


filtros = [
    # estabelecimento
    {"posicao": (0, 2), "valor": "57"},
    {"posicao": (2, 4), "valor": "00"},
    {"posicao": (166, 167), "valor": "1"}
    # principal
        # {"posicao": (0, 2), "valor": "57"},
        # {"posicao": (2, 4), "valor": "00"},
        # {"posicao": (166, 167), "valor": "2"}
    # detalhamento elo
        # {"posicao": (0, 2), "valor": "57"},
        # {"posicao": (2, 4), "valor": "01"},
        # {"posicao": (4, 6), "valor": "EL"}
    # detalhamento visa
       # {"posicao": (0, 2), "valor": "57"},
       # {"posicao": (2, 4), "valor": "0D"},
    # detalhamento master
       # {"posicao": (0, 2), "valor": "57"},
       # {"posicao": (2, 4), "valor": "05"},
]

eh_diretorio = os.path.isdir(alvo)
contador = Counter()

try:
    if eh_diretorio:
        alvos = [os.path.join(alvo, f) for f in os.listdir(alvo) if os.path.isfile(os.path.join(alvo, f))]
    else:
        alvos = [alvo]
    
    for arq in alvos:
        with open(arq, 'r') as f:
            for linha in f:
                # Aplicar filtros
                if any(linha[pos[0]:pos[1]] != filtro["valor"] for filtro in filtros for pos in [filtro["posicao"]]):
                    continue
                if lista:
                    print(linha.strip())
                    continue
                valor = linha[coluna_inicio:coluna_fim].strip()
                contador[valor] += 1
    if not lista:
        print(f"\nOcorrências posicao coluna [{coluna_inicio}:{coluna_fim}]:")
        for valor, count in sorted(contador.items()):
            print(f"{valor}: {count}")
except FileNotFoundError:
    print(f"Erro: alvo '{alvo}' não encontrado")
except ValueError:
    print("Erro: número da coluna deve ser um inteiro")