import os
def get_env_var(var_name):
    var_value = os.getenv(var_name)
    if var_value is None:
        raise ValueError(f"Environment variable {var_name} not set.")
    return var_value
