from flask import Flask
from app.db import db
from app.routes import book_api as api_blueprint
from app.models import Book

def create_app():
    app = Flask(__name__)
    
    # Configuration
    app.config['SQLALCHEMY_DATABASE_URI'] = 'sqlite:///app.db'
    app.config['SQLALCHEMY_TRACK_MODIFICATIONS'] = False

    # Initialize extensions
    db.init_app(app)
    
    # Register blueprints
    app.register_blueprint(api_blueprint)
    
    # Ensure tables are created without migrations
    with app.app_context():
        db.create_all()

    return app
