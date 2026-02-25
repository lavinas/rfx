import sys
from collections import Counter
import os

if len(sys.argv) < 3:
    print("Uso: python count_dimp.py <arquivo.txt> <número_coluna>")
    sys.exit(1)

arquivo = sys.argv[1]
coluna = int(sys.argv[2])
filtrar = sys.argv[3] if len(sys.argv) > 3 else None
eh_diretorio = os.path.isdir(arquivo)
contador = Counter()

try:
    if eh_diretorio:
        arquivos = [os.path.join(arquivo, f) for f in os.listdir(arquivo) if os.path.isfile(os.path.join(arquivo, f))]
    else:
        arquivos = [arquivo]
    
    for arq in arquivos:
        with open(arq, 'r', encoding='iso-8859-1') as f:
            for linha in f:
                partes = linha.strip().split('|')
                if filtrar is not None and partes[1] != filtrar:
                    continue
                if coluna < len(partes):
                    valor = partes[coluna].strip()
                    contador[valor] += 1
    
    print(f"\nOcorrências da coluna {coluna}:")
    for valor, count in sorted(contador.items()):
        print(f"{valor}: {count}")
except FileNotFoundError:
    print(f"Erro: arquivo '{arquivo}' não encontrado")
except ValueError:
    print("Erro: número da coluna deve ser um inteiro")