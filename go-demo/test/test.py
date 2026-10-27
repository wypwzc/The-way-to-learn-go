from selenium import webdriver
from selenium.webdriver.chrome.options import Options

options = Options()
options.add_experimental_option("detach", True)  # 关键：让浏览器脱离脚本生命周期

driver = webdriver.Chrome(options=options)
driver.get("https://wzcwyp.xyz:5888/admin")