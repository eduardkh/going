import subprocess

proc = subprocess.Popen(
    ["./whois_isracard.co.il_validity.sh"], stdout=subprocess.PIPE, shell=True)
(out, err) = proc.communicate()
shell = out.decode("utf-8")
print(f"This is an output from the WHOIS command saved as a shell script and initiated by Python: {shell}")
