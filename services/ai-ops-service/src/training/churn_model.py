import pandas as pd
from sklearn.model_selection import train_test_split
from sklearn.ensemble import RandomForestClassifier
import joblib

class ChurnTrainingPipeline:
    def train(self, data_path: str, model_output_path: str):
        """
        Trains a Churn prediction model using historical usage and billing data.
        """
        # 1. Load Data
        df = pd.read_csv(data_path)

        # 2. Preprocess (Feature Engineering)
        X = df.drop('churn', axis=1)
        y = df['churn']

        X_train, X_test, y_train, y_test = train_test_split(X, y, test_size=0.2)

        # 3. Train Model
        model = RandomForestClassifier(n_estimators=100)
        model.fit(X_train, y_train)

        # 4. Save Model
        joblib.dump(model, model_output_path)
        print(f"Model saved to {model_output_path}")

if __name__ == "__main__":
    # pipeline = ChurnTrainingPipeline()
    # pipeline.train('data/churn_data.csv', 'models/churn_v1.joblib')
    pass
