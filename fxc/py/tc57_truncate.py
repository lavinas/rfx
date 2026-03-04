import sys
import os
from pathlib import Path

def process_files(directory_in, directory_out):
    if not os.path.isdir(directory_in):
        print(f"Erro: {directory_in} não é um diretório válido")
        return

    if not os.path.exists(directory_out):
        os.makedirs(directory_out)

    for filename in os.listdir(directory_in):
        filepath = os.path.join(directory_in, filename)
        
        if not os.path.isfile(filepath):
            continue
        
        output_path = os.path.join(directory_out, filename)
        
        try:
            with open(filepath, 'r', encoding='utf-8', errors='ignore') as infile:
                with open(output_path, 'w', encoding='utf-8') as outfile:
                    for line in infile:
                        if line.startswith("5700") and len(line) > 167 and line[166] == '2':
                            if len(line) >= 68:
                                modified_line = line[:57] + "*" * 11 + line[68:]
                                outfile.write(modified_line)
                            else:
                                outfile.write(line)
                        else:
                            outfile.write(line)
            print(f"Processado: {filepath} -> {output_path}")
        except Exception as e:
            print(f"Erro ao processar {filepath}: {e}")

if __name__ == "__main__":
    if len(sys.argv) < 2:
        print("Uso: python truncate_tc57.py <diretório>")
        sys.exit(1)
    
    directory_in = sys.argv[1]
    directory_out = sys.argv[2] if len(sys.argv) > 2 else directory_in
    process_files(directory_in, directory_out)