import yaml
from ezored.models.dependency import Dependency

from .constants import Constants
from .logger import Logger
from .target import Target


class Project(object):
    config = {}
    dependencies = []
    targets = []

    @staticmethod
    def create_from_project_file():
        Logger.d('Loading project from file...')

        try:
            with open(Constants.PROJECT_FILE, 'r') as stream:
                return Project.create_from_stream(stream)
        except IOError as exc:
            Logger.f('Error while read project file: {0}'.format(exc))

    @staticmethod
    def create_from_stream(stream):
        try:
            # load project data
            project = Project()

            file_data = yaml.load(stream)

            # configurations
            Logger.d('Loading project configurations...')

            if 'config' in file_data:
                config = file_data['config']

                if type(config) is dict:
                    project.config = config

            # dependencies
            Logger.d('Loading project dependencies...')

            if 'dependencies' in file_data:
                dependencies = file_data['dependencies']

                for dependency_data in dependencies:
                    dependency = Dependency.from_dict(dependency_data)
                    project.dependencies.append(dependency)

            # targets
            Logger.d('Loading project targets...')

            if 'targets' in file_data:
                targets = file_data['targets']

                for target_data in targets:
                    target = Target.from_dict(target_data)
                    project.targets.append(target)

            # finished
            Logger.d('Project loaded')

            return project
        except yaml.YAMLError as exc:
            Logger.f('Error while load project data: {0}'.format(exc))

    def get_config_value(self, key):
        if self.config and key in self.config:
            return self.config[key]

        return None
